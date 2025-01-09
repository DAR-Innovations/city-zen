from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from typing import List, Optional

from app.models.events.issue import Issues
from app.modules.issues.types.issues_dtos import PostIssueRequestDTO, IssueResponseDTO


class IssuesService:
    def __init__(self, db: AsyncSession):
        self.db = db

    async def create_issue(
            self,
            issue_data: PostIssueRequestDTO,
            author_id: int
    ) -> IssueResponseDTO:
        issue = Issues(
            name=issue_data.name,
            description=issue_data.description,
            longitude=issue_data.longitude,
            latitude=issue_data.latitude,
            author_id=author_id
        )

        self.db.add(issue)
        await self.db.flush()
        await self.db.commit()
        await self.db.refresh(issue)

        return self._to_response_dto(issue)

    async def get_issues_by_author(
            self,
            author_id: int
    ) -> List[IssueResponseDTO]:
        query = select(Issues).where(
            Issues.author_id == author_id,
            Issues.deleted_at.is_(None)
        )
        result = await self.db.execute(query)
        issues = result.scalars().all()

        return [self._to_response_dto(issue) for issue in issues]

    async def get_issue_by_id(
            self,
            issue_id: int
    ) -> Optional[IssueResponseDTO]:
        query = select(Issues).where(
            Issues.id == issue_id,
            Issues.deleted_at.is_(None)
        )
        result = await self.db.execute(query)
        issue = result.scalars().first()

        if not issue:
            return None

        return self._to_response_dto(issue)

    def _to_response_dto(self, issue: Issues) -> IssueResponseDTO:
        return IssueResponseDTO(
            id=issue.id,
            name=issue.name,
            description=issue.description,
            longitude=float(issue.longitude),
            latitude=float(issue.latitude),
            is_completed=issue.is_completed,
            created_at=issue.created_at.isoformat(),
            image_url=issue.image_url or "",
            author_id=issue.author_id
        )
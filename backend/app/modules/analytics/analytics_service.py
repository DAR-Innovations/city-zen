from sqlalchemy.orm import Session
from typing import List

from app.db.database import get_db
from fastapi import Depends
from sqlalchemy import func, Float

from app.models.events.issue import Issues
from app.models.tasks.department_tasks import DepartmentTasks
from app.models.tasks.volunteer_tasks import VolunteerTasks
from app.modules.analytics.types.analytics_dtos import RegionTasksResponseDTO, CompletedTasksDTO, RegionTasksRequestDTO

EARTH_RADIUS_KM = 6371.0

def haversine(lat1, lon1, lat2, lon2):
    return func.acos(
        func.sin(func.radians(lat1)) * func.sin(func.radians(lat2)) +
        func.cos(func.radians(lat1)) * func.cos(func.radians(lat2)) *
        func.cos(func.radians(lon2) - func.radians(lon1))
    ) * EARTH_RADIUS_KM

class AnalyticsService:
    def __init__(self, db: Session = Depends(get_db)):
        self.db = db

    async def get_completed_tasks_by_department(self, department_id: int) -> CompletedTasksDTO:
        completed_tasks = self.db.query(DepartmentTasks).filter(
            DepartmentTasks.department_id == department_id,
            DepartmentTasks.status == "COMPLETED"
        ).count()

        return CompletedTasksDTO(completedTasks=completed_tasks)

    async def get_volunteers_completed_tasks(self, volunteer_id: int) -> CompletedTasksDTO:
        completed_tasks = self.db.query(VolunteerTasks).filter(
            VolunteerTasks.volunteer_id == volunteer_id,
            VolunteerTasks.status == "COMPLETED"
        ).count()

        return CompletedTasksDTO(completedTasks=completed_tasks)

    async def get_analytics_in_area(self, dto: RegionTasksRequestDTO) -> RegionTasksResponseDTO:
        # Extracting the center and radius from the request DTO
        center_lat = dto.latitude
        center_lon = dto.longitude
        radius = dto.radius

        # Performing the query with Haversine formula
        query = self.db.query(
            Issues,
            haversine(Issues.latitude, Issues.longitude, center_lat, center_lon).label('distance')
        ).filter(
            haversine(Issues.latitude, Issues.longitude, center_lat, center_lon) <= radius
        )

        # Counting the created issues and resolved tasks
        created_issues_count = 0
        resolved_issues_count = 0

        for issue, _ in query.all():  # Note: we are ignoring the distance value
            created_issues_count += 1
            if issue.is_completed:
                resolved_issues_count += 1


        return RegionTasksResponseDTO(
            createdIssues=created_issues_count,
            resolvedTasks=resolved_issues_count
        )
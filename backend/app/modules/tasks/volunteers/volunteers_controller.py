from fastapi import APIRouter, Depends, UploadFile, HTTPException
from sqlalchemy.orm import Session
from app.db.database import get_db
from app.models.tasks.volunteer_tasks import VolunteerTasks
from app.models.events.reports import Reports
from app.modules.tasks.common.reported_by import resolve_reported_by
from app.ai.validate_task import validate_task_completion

router = APIRouter()

@router.get("/tasks")
def get_volunteer_tasks(
    urgency: str = None,
    complexity: str = None,
    db: Session = Depends(get_db)
):
    query = db.query(VolunteerTasks).filter(VolunteerTasks.status == "PENDING")
    if urgency:
        query = query.filter(VolunteerTasks.urgency == urgency)
    if complexity:
        query = query.filter(VolunteerTasks.complexity == complexity)
    return {"tasks": query.all()}

@router.post("/tasks/{task_id}/report")
async def complete_task(
    task_id: int,
    file: UploadFile,
    comment: str,
    reported_by_id: int,
    db: Session = Depends(get_db)
):
    task = db.query(VolunteerTasks).filter(VolunteerTasks.id == task_id).first()
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    reporter = resolve_reported_by(db, reported_by_id)
    is_valid = await validate_task_completion(file, task)
    if is_valid:
        task.status = "COMPLETED"
        report = Reports(issue_id=task.issue_id, description=comment, reported_by=reported_by_id)
        db.add(report)
        db.commit()
        return {"message": "Task completed successfully"}
    return {"message": "Task validation failed"}

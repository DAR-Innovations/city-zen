from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from app.db.database import get_db
from app.models.tasks.department_tasks import DepartmentTasks

router = APIRouter()

@router.get("/tasks")
def get_department_tasks(
    status: str = None,
    urgency: str = None,
    db: Session = Depends(get_db)
):
    query = db.query(DepartmentTasks).filter(DepartmentTasks.status == "PENDING")
    if status:
        query = query.filter(DepartmentTasks.status == status)
    if urgency:
        query = query.filter(DepartmentTasks.urgency == urgency)
    return {"tasks": query.all()}

@router.post("/tasks/{task_id}/assign")
def assign_task(task_id: int, employee_id: int, db: Session = Depends(get_db)):
    task = db.query(DepartmentTasks).filter(DepartmentTasks.id == task_id).first()
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    task.assigned_to = employee_id
    db.commit()
    return {"message": "Task assigned"}

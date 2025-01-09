from fastapi import HTTPException
from sqlalchemy.orm import Session
from app.models.agents.users import Users
from app.models.agents.employees import Employees

def resolve_reported_by(db: Session, reported_by_id: int):
    """
    Resolve whether the reporter is a User or an Employee.
    :param db: Database session.
    :param reported_by_id: ID of the reporter.
    :return: A dictionary with the type ("user" or "employee") and the entity object.
    """
    user = db.query(Users).filter(Users.id == reported_by_id).first()
    if user:
        return {"type": "user", "entity": user}

    employee = db.query(Employees).filter(Employees.id == reported_by_id).first()
    if employee:
        return {"type": "employee", "entity": employee}

    raise HTTPException(status_code=404, detail="Reporter not found")

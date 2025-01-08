from sqlalchemy import Column, String, Text, Integer, ForeignKey
from app.models.base import Base

class DepartmentTask(Base):
    issue_id = Column(Integer, ForeignKey("issue.id"))
    title = Column(String(255), nullable=False)
    description = Column(Text)
    task_type_id = Column(Integer, ForeignKey("tasktypes.id"))
    status = Column(String(50), nullable=False, default="PENDING")
    urgency = Column(String(50))
    complexity = Column(String(50))
    department_id = Column(Integer, ForeignKey("departments.id"))

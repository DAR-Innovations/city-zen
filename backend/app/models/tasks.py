from sqlalchemy import Integer, Column, String, Date, ForeignKey
from sqlalchemy.orm import relationship
from app.models.base import Base

class Task(Base):
    issue_id = Column(Integer, ForeignKey("issue.id"))
    task_type_id = Column(Integer, ForeignKey("task_types.id"))
    urgency = Column(String(50), nullable=False)
    complexity = Column(String(50), nullable=False)
    processed_at = Column(Date)
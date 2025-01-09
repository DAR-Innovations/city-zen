from sqlalchemy import Column, String, Text, Integer, ForeignKey
from app.models.base import Base

class VolunteerTasks(Base):
    issue_id = Column(Integer, ForeignKey("issues.id"))
    title = Column(String(255), nullable=False)
    description = Column(Text)
    status = Column(String(50), nullable=False, default="PENDING")
    urgency = Column(String(50))
    complexity = Column(String(50))
    volunteer_id = Column(Integer, ForeignKey("users.id"), nullable=True)

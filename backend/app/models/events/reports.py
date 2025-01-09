from sqlalchemy import Integer, Column, String, Text, ForeignKey, DateTime
from sqlalchemy.orm import relationship
from datetime import datetime
from app.models.base import Base

class BaseReport(Base):
    __abstract__ = True
    issue_id = Column(Integer, ForeignKey("issues.id"))
    description = Column(Text)
    image_url = Column(String(255))
    reported_by = Column(Integer)
    reported_at = Column(DateTime, default=datetime.now)

class UserReports(BaseReport):
    reported_by = Column(Integer, ForeignKey("users.id"))

class DepartmentReports(BaseReport):
    reported_by = Column(Integer, ForeignKey("departments.id"))
from sqlalchemy import Integer, Column, String, Text, ForeignKey
from sqlalchemy.orm import relationship
from app.models.base import Base

class Report(Base):
    issue_id = Column(Integer, ForeignKey("issue.id"))
    description = Column(Text)
    image = Column(String(255))
    department_id = Column(Integer, ForeignKey("departments.id"))
    user_id = Column(Integer, ForeignKey("users.id"))

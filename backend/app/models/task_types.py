from sqlalchemy import Integer, Column, String, Text, ForeignKey
from sqlalchemy.orm import relationship
from app.models.base import Base

class TaskType(Base):
    name = Column(String(255), nullable=False)
    description = Column(Text)
    department_id = Column(Integer, ForeignKey("departments.id"))
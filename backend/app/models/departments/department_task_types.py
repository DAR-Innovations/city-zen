from sqlalchemy import Column, Integer, ForeignKey
from app.models.base import Base

class DepartmentTaskType(Base):
    department_id = Column(Integer, ForeignKey("departments.id"))
    task_type_id = Column(Integer, ForeignKey("tasktypes.id"))

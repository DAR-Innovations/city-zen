from sqlalchemy import Column, Integer, String, Boolean, ForeignKey
from app.models.base import Base

class Employees(Base):
    name = Column(String(255), nullable=False)
    phone = Column(String(20), unique=True, nullable=False)
    is_verified = Column(Boolean, default=False)
    department_id = Column(Integer, ForeignKey("departments.id"))
    role = Column(String(50), nullable=False)
    password = Column(String(255), nullable=False)

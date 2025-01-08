from sqlalchemy import Column, String, Text
from app.models.base import Base

class Department(Base):
    name = Column(String(255), nullable=False)
    description = Column(Text)
    phone = Column(String(20))

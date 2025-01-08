from sqlalchemy import Column, String, Boolean
from app.models.base import Base

class User(Base):
    name = Column(String(255), nullable=False)
    phone = Column(String(20), unique=True, nullable=False)
    is_verified = Column(Boolean, default=False)
    role = Column(String(50), nullable=False)
    password = Column(String(255), nullable=False)

from sqlalchemy import Integer, Column, String, Text, Boolean, ForeignKey, DECIMAL
from sqlalchemy.orm import relationship
from app.models.base import Base


class Issue(Base):
    name = Column(String(255), nullable=False)
    description = Column(Text, nullable=False)
    tags = Column(String(255))
    status = Column(String(50), nullable=False, default='PENDING')
    is_public = Column(Boolean, default=True)
    longitude = Column(DECIMAL(9, 6))
    latitude = Column(DECIMAL(9, 6))
    user_id = Column(Integer, ForeignKey("users.id"))
    image = Column(String(255))

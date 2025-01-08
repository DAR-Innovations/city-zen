from sqlalchemy import Integer, Column, String, Text, Boolean, ForeignKey, DECIMAL
from sqlalchemy.orm import relationship
from app.models.base import Base

class Issue(Base):
    name = Column(String(255), nullable=False)
    description = Column(Text, nullable=False)
    is_completed = Column(Boolean, default=False)
    longitude = Column(DECIMAL(9, 6))
    latitude = Column(DECIMAL(9, 6))
    author_id = Column(Integer, ForeignKey("users.id"))
    image_url = Column(String(255))

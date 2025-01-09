from sqlalchemy import Integer, Column, String, Text, ForeignKey, DateTime
from sqlalchemy.orm import relationship
from datetime import datetime
from app.models.base import Base

class Reports(Base):
    issue_id = Column(Integer, ForeignKey("issues.id"))
    description = Column(Text)
    image_url = Column(String(255))
    reported_by = Column(Integer, ForeignKey("users.id"))
    reported_at = Column(DateTime, default=datetime.now)

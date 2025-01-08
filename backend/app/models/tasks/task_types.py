from app.models.base import Base
from sqlalchemy import Column, String, Text

class TaskTypes(Base):
    name = Column(String(255), nullable=False)
    description = Column(Text)

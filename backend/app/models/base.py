from sqlalchemy.ext.declarative import as_declarative, declared_attr
from sqlalchemy import Column, Integer, DateTime
from datetime import datetime
import re

@as_declarative()
class Base:
    id = Column(Integer, primary_key=True, index=True)
    created_at = Column(DateTime, default=datetime.now, nullable=False)                         # Automatically set at creation
    updated_at = Column(DateTime, default=datetime.now, onupdate=datetime.now, nullable=False)  # Automatically updated
    deleted_at = Column(DateTime, nullable=True)                                                # For soft deletes

    @declared_attr
    def __tablename__(cls):
        name = cls.__name__
        return re.sub(r'(?<!^)(?=[A-Z])', '_', name).lower()

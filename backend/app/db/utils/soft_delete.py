from datetime import datetime
from sqlalchemy.orm import Session

def soft_delete(record, db: Session):
    """
    Marks a record as deleted by setting its `deleted_at` timestamp.
    """
    record.deleted_at = datetime.now()
    db.commit()
    db.refresh(record)

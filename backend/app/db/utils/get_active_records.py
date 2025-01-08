from sqlalchemy.orm import Query

def get_active_records(query: Query):
    """
    Filters out soft-deleted records.
    """
    return query.filter_by(deleted_at=None)

import psycopg2
from psycopg2.extras import RealDictCursor
from dotenv import load_dotenv
import os

load_dotenv()

DB_CONFIG = {
    "host": os.getenv("DB_HOST", "localhost"),
    "port": int(os.getenv("DB_PORT", 5432)),
    "database": os.getenv("DB_NAME", "city-zen-db"),
    "user": os.getenv("DB_USER", "postgres"),
    "password": os.getenv("DB_PASSWORD", ""),
}

def get_db_connection():
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        return conn
    except Exception as e:
        raise RuntimeError(f"Database connection failed: {e}")
    

def get_departments():
    """
    Fetch the list of departments from the database.
    Example: ['Police', 'Public Works']
    """
    conn = get_db_connection()
    try:
        with conn.cursor(cursor_factory=RealDictCursor) as cursor:
            query = "SELECT id, name FROM departments"
            cursor.execute(query)
            rows = cursor.fetchall()
            return [{"id": row["id"], "name": row["name"]} for row in rows]
    except Exception as e:
        raise RuntimeError(f"Failed to fetch departments: {e}")
    finally:
        conn.close()

def get_task_types_by_department(department_id):
    """
    Fetch task types associated with a specific department.
    Example: ['Theft', 'Vandalism']
    """
    conn = get_db_connection()
    try:
        with conn.cursor(cursor_factory=RealDictCursor) as cursor:
            query = """
            SELECT tt.name AS task_type_name
            FROM department_task_types dtt
            JOIN task_types tt ON dtt.task_type_id = tt.id
            WHERE dtt.department_id = %s
            """
            cursor.execute(query, (department_id,))
            rows = cursor.fetchall()
            return [row["task_type_name"] for row in rows]
    except Exception as e:
        raise RuntimeError(f"Failed to fetch task types for department {department_id}: {e}")
    finally:
        conn.close()


def fetch_departments_and_task_types():
    """
    Fetch all departments and their associated task types.
    Example: {'Police': ['Theft', 'Vandalism'], 'Public Works': ['Road Repair']}
    """
    departments = get_departments()
    department_tasks = {}

    for department in departments:
        task_types = get_task_types_by_department(department["id"])
        department_tasks[department["name"]] = task_types

    return department_tasks
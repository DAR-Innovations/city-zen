from fastapi import FastAPI
from app.db.database import engine, get_db
from app.models.base import Base
from app.router import api_router
import uvicorn

def init_db():
    """
    Initialize the database by creating all tables.
    This function should only be used in development or testing.
    """
    Base.metadata.create_all(bind=engine)

from fastapi import FastAPI
from contextlib import asynccontextmanager
from app.db.database import engine
from app.models.base import Base
from app.router import api_router

from app.models.departments.departments import Departments
from app.models.agents.employees import Employees
from app.models.agents.users import Users
from app.models.tasks.task_types import TaskTypes
from app.models.departments.department_task_types import DepartmentTaskTypes
from app.models.events.issue import Issues
from app.models.events.reports import Reports
from app.models.tasks.volunteer_tasks import VolunteerTasks
from app.models.tasks.department_tasks import DepartmentTasks

@asynccontextmanager
async def lifespan(app: FastAPI):
    print("Starting the application...")
    print("Initializing the database...")
    Base.metadata.create_all(bind=engine)
    print("Database initialized successfully!")
    yield
    print("Shutting down the application...")

app = FastAPI(lifespan=lifespan)

app.include_router(api_router)

@app.get("/")
def root():
    return {"message": "Welcome to the Citizen Problem Reporting API!"}


if __name__ == "__main__":
    uvicorn.run("main:app", host="127.0.0.1", port=8000, reload=True)
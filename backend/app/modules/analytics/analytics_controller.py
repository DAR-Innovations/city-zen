from fastapi import APIRouter, Depends, HTTPException
from typing import List

from app.modules.analytics.analytics_service import AnalyticsService
from app.modules.analytics.types.analytics_dtos import RegionTasksResponseDTO, CompletedTasksDTO, RegionTasksRequestDTO

router = APIRouter()

@router.get("/department/{department_id}", response_model=CompletedTasksDTO)
async def get_completed_tasks_by_department(
    department_id: int,
    service: AnalyticsService = Depends()
):
    try:
        return await service.get_completed_tasks_by_department(department_id)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/volunteer/{volunteer_id}", response_model=CompletedTasksDTO)
async def get_volunteers_completed_tasks(
    volunteer_id: int,
    service: AnalyticsService = Depends()
):
    try:
        return await service.get_volunteers_completed_tasks(volunteer_id)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/area", response_model=RegionTasksResponseDTO)
async def get_analytics_in_area(
    region: RegionTasksRequestDTO,
    service: AnalyticsService = Depends()
):
    try:
        return await service.get_analytics_in_area(region)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
from fastapi import APIRouter, Depends, HTTPException, Request, File, UploadFile, Form
from sqlalchemy.orm import Session
from typing import List
import json

from app.db.database import get_db

from app.modules.images.images_service import save_image_with_metadata
from app.modules.issues.issues_service import IssuesService
from app.modules.issues.types.issues_dtos import PostIssueRequestDTO, IssueResponseDTO

router = APIRouter()

@router.post('/issues', response_model=IssueResponseDTO)
async def create_issue(
        request: Request,
        name: str = Form(...),
        description: str = Form(...),
        longitude: float = Form(...),
        latitude: float = Form(...),
        image: UploadFile = File(...),
        db: Session = Depends(get_db)
):
    try:
        valid_types = ["image/jpeg", "image/png"]
        if image.content_type not in valid_types:
            raise HTTPException(
                status_code=400,
                detail="Invalid file type. Only JPEG and PNG are allowed."
            )

        file_path, metadata_path = await save_image_with_metadata(image)

        issue_data = PostIssueRequestDTO(
            name=name,
            description=description,
            longitude=longitude,
            latitude=latitude
        )

        image_coordinates = await extract_coordinates_from_image(file_path)

        final_coordinates = {
            'longitude': image_coordinates['longitude'] if image_coordinates else issue_data.longitude,
            'latitude': image_coordinates['latitude'] if image_coordinates else issue_data.latitude
        }

        issue_service = IssuesService(db)
        return await issue_service.create_issue(
            issue_data=issue_data,
            author_id=request.state.user.id,
            image_path=file_path,
            longitude=final_coordinates['longitude'],
            latitude=final_coordinates['latitude']
        )
    except HTTPException as he:
        raise he
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.get('/issues/my', response_model=List[IssueResponseDTO])
async def get_my_issues(
        request: Request,
        db: Session = Depends(get_db)
):
    try:
        issue_service = IssuesService(db)
        return await issue_service.get_issues_by_author(
            author_id=request.state.user.id
        )
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.get('/issues/{issue_id}', response_model=IssueResponseDTO)
async def get_issue_by_id(
        issue_id: int,
        db: Session = Depends(get_db)
):
    try:
        issue_service = IssuesService(db)
        issue = await issue_service.get_issue_by_id(issue_id=issue_id)
        if not issue:
            raise HTTPException(status_code=404, detail="Issue not found")
        return issue
    except HTTPException as he:
        raise he
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


async def extract_coordinates_from_image(file_path: str) -> dict:
    try:
        from PIL import Image, ExifTags
        image = Image.open(file_path)
        exif_data = image._getexif()

        if not exif_data:
            return None

        # Find the GPSInfo tag
        gps_info = None
        for tag_id in ExifTags.TAGS:
            if ExifTags.TAGS[tag_id] == 'GPSInfo':
                if tag_id in exif_data:
                    gps_info = exif_data[tag_id]
                break

        if not gps_info:
            return None

        def convert_to_degrees(value):
            d, m, s = value
            return d + (m / 60.0) + (s / 3600.0)

        latitude = convert_to_degrees(gps_info[2])
        longitude = convert_to_degrees(gps_info[4])

        # If south, make latitude negative
        if gps_info[1] == 'S':
            latitude = -latitude

        # If west, make longitude negative
        if gps_info[3] == 'W':
            longitude = -longitude

        return {
            'latitude': latitude,
            'longitude': longitude
        }
    except Exception:
        return None
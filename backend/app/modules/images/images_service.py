import os
from uuid import uuid4
from datetime import datetime
from PIL import Image, ExifTags
from fastapi import UploadFile
from fastapi.exceptions import HTTPException
from app.core.config import settings

UPLOAD_FOLDER = settings.UPLOAD_FOLDER

if not os.path.exists(UPLOAD_FOLDER):
    os.makedirs(UPLOAD_FOLDER)

def extract_exif_data(image_path: str) -> dict:
    """
    Extract EXIF metadata from an image.
    :param image_path: Path to the image file.
    :return: Dictionary containing EXIF metadata.
    """
    try:
        image = Image.open(image_path)
        exif_data = image._getexif() or {}
        exif = {
            ExifTags.TAGS[k]: v
            for k, v in exif_data.items()
            if k in ExifTags.TAGS
        }
   
        gps_info = exif.get("GPSInfo")
        if gps_info:
            gps_data = {
                ExifTags.GPSTAGS.get(k, k): v
                for k, v in gps_info.items()
            }
            exif["GPSInfo"] = gps_data

        return exif
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Failed to extract EXIF data: {str(e)}")

async def save_image_with_metadata(file: UploadFile) -> tuple[str, str]:
    """
    Save an uploaded image file and its detailed metadata in the uploads directory.
    :param file: The uploaded file.
    :return: Tuple containing the image file path and metadata file path.
    """
    try:
        
        file_extension = file.filename.split(".")[-1]
        unique_filename = f"{uuid4().hex}.{file_extension}"

        folder_name = uuid4().hex
        folder_path = os.path.join(UPLOAD_FOLDER, folder_name)
        os.makedirs(folder_path, exist_ok=True)
        
        file_path = os.path.join(folder_path, unique_filename)
        with open(file_path, "wb") as f:
            content = await file.read()
            f.write(content)

        exif_data = extract_exif_data(file_path)

        metadata = {
            "original_filename": file.filename,
            "content_type": file.content_type,
            "file_size": len(content),
            "uploaded_at": datetime.utcnow().isoformat(),
            "exif_data": exif_data,
        }

        metadata_path = os.path.join(folder_path, f"{unique_filename}.txt")
        with open(metadata_path, "w") as f:
            for key, value in metadata.items():
                if isinstance(value, dict):
                    f.write(f"{key}:\n")
                    for sub_key, sub_value in value.items():
                        f.write(f"  {sub_key}: {sub_value}\n")
                else:
                    f.write(f"{key}: {value}\n")

        return file_path, metadata_path

    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Failed to save image or metadata: {str(e)}")

from pydantic import BaseModel, EmailStr, Field, field_validator
from datetime import date

class CreateUserValidation(BaseModel):
    nama: str = Field(..., max_length=100, description="Nama harus berupa teks")
    email: EmailStr = Field(..., description="Email tidak valid")
    password: str = Field(..., min_length=6, description="Password minimal harus memiliki 6 karakter")
    NIK: str = Field(..., pattern=r"^\d{16}$", description="NIK harus terdiri dari 16 karakter dan hanya angka")
    alamat: str = Field(..., max_length=100, description="Alamat tidak boleh lebih dari 100 karakter")
    telepon: str = Field(..., pattern=r"^\d{1,15}$", description="Nomor telepon harus terdiri dari angka maksimal 15 karakter")
    jenis_kelamin: str = Field(..., pattern=r"^(L|P)$", description="Jenis kelamin harus 'L' (Laki-laki) atau 'P' (Perempuan)")
    kepala_keluarga: int = Field(..., description="Kepala keluarga harus bernilai 0 atau 1")
    tempat_lahir: str = Field(..., max_length=50, description="Tempat lahir tidak boleh lebih dari 50 karakter")
    tanggal_lahir: date = Field(..., description="Tanggal lahir harus dalam format yang valid (YYYY-MM-DD)")
    jenis_usaha: str = Field(..., max_length=50, description="Jenis usaha tidak boleh lebih dari 50 karakter")

    @field_validator('kepala_keluarga')
    def validate_kepala_keluarga(cls, value):
        if value not in [0, 1]:
            raise ValueError('Kepala keluarga harus bernilai 0 atau 1')
        return value

class LoginUserValidation(BaseModel):
    email: EmailStr = Field(..., description="Email tidak valid")
    password: str = Field(..., min_length=6, description="Password minimal harus memiliki 6 karakter")

class UpdateUserValidation(BaseModel):
    nama: str = Field(None, max_length=100, description="Nama harus berupa teks dan maksimal 100 karakter")
    email: EmailStr = Field(None, description="Email tidak valid")
    password: str = Field(None, min_length=6, description="Password minimal harus memiliki 6 karakter")
    NIK: str = Field(None, pattern=r"^\d{16}$", description="NIK harus terdiri dari 16 karakter dan hanya angka")
    alamat: str = Field(None, max_length=100, description="Alamat tidak boleh lebih dari 100 karakter")
    telepon: str = Field(None, pattern=r"^\d{1,15}$", description="Nomor telepon harus terdiri dari angka maksimal 15 karakter")
    jenis_kelamin: str = Field(None, pattern=r"^(L|P)$", description="Jenis kelamin harus 'L' (Laki-laki) atau 'P' (Perempuan)")
    kepala_keluarga: int = Field(None, description="Kepala keluarga harus bernilai 0 atau 1")
    tempat_lahir: str = Field(None, max_length=50, description="Tempat lahir tidak boleh lebih dari 50 karakter")
    tanggal_lahir: date = Field(None, description="Tanggal lahir harus dalam format yang valid (YYYY-MM-DD)")
    jenis_usaha: str = Field(None, max_length=50, description="Jenis usaha tidak boleh lebih dari 50 karakter")

    @field_validator('kepala_keluarga')
    def validate_kepala_keluarga(cls, value):
        if value is not None and value not in [0, 1]:
            raise ValueError('Kepala keluarga harus bernilai 0 atau 1')
        return value
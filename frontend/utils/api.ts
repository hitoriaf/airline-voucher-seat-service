import axios from "axios";

const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL;

const api = axios.create({
  baseURL: baseURL,
  headers: {
    "Content-Type": "application/json",
  },    
});

type CheckRequestDto = {
    flightNumber: string;
    date: string;
}

type CheckResponseDto = {
    exist: boolean;
}

type GenerateRequestDto = {
    name: string;
    id: string;
    flightNumber: string;
    date: string;
    aircraft: string;
}

type GenerateResponseDto = {
    success: boolean;
    seats: string[];
}

export async function checkVoucher(payload: CheckRequestDto): Promise<CheckResponseDto> {
    const response = await api.post<CheckResponseDto>("/api/check", payload);
    return response.data;
}

export async function generateVoucher(payload: GenerateRequestDto): Promise<GenerateResponseDto> {
    const response = await api.post<GenerateResponseDto>("/api/generate", payload);
    return response.data;
}
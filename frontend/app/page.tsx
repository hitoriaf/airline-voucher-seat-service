"use client";

import { useState } from 'react';
import { checkVoucher, generateVoucher } from '@/utils/api';
const Page = () => {
  const [formData, setFormData] = useState({ 
    name: '',
    id: '',
    flightNumber: '',
    date: '',
    aircraft: '',
  });
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<{
    success: boolean;
    seats: string[];
  } | null>(null);
  const [checkResult, setCheckResult] = useState<{
    exist: boolean;
  } | null>(null);
  const aircraft = [
    "ATR", 
    "Airbus 320", 
    "Boeing 737 Max"
  ]
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setResult(null);
    setCheckResult(null);
    try {
      const response = await checkVoucher({
        flightNumber: formData.flightNumber,
        date: formData.date
      });
      if(response.exist) {
        setResult(null);
        setCheckResult(response);
        // alert("Vouchers already generated for this flight.");
        return;
      }
      console.log("Form Data Submitted:", formData);
      const generateResponse = await generateVoucher(formData);
      if(generateResponse.success) {
        setResult(generateResponse);
      } else {
        alert("Failed to generate vouchers. Please try again.");
      }
    } catch (error) {
      console.error("Error during voucher process:", error);
      alert("An error occurred. Please try again.");
    } finally {
      setLoading(false);
    }
  }
  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { id, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [id]: value
    }));
  }
  return <div className="container mx-auto px-4 h-screen flex items-center justify-center">
    <div className="flex flex-col gap-6 w-full bg-[#d5e8fa] p-6 rounded-lg shadow-md">
      <h1 className="text-4xl font-bold text-center">Welcome to the Airline Voucher Seat Service</h1>
      <div className="p-6 w-full max-w-lg mx-auto">
        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor="name" className="block text-lg font-medium text-gray-700 mb-2">Crew Name</label>
            <input 
              required
              type="text" 
              id="name" 
              className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-4 px-4 bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500" 
              placeholder="Enter Crew Name"
              value ={formData.name}
              onChange={handleChange}
            />
          </div>
          <div>
            <label htmlFor="id" className="block text-lg font-medium text-gray-700 mb-2 mt-4">Crew ID</label>
            <input 
              required
              type="text" 
              id="id" 
              className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-4 px-4 bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500" 
              placeholder="Enter Crew ID"
              value ={formData.id}
              onChange={handleChange}
            />
          </div>
          <div>
            <label htmlFor="flightNumber" className="block text-lg font-medium text-gray-700 mb-2 mt-4">Flight Number</label>
            <input 
              required
              type="text" 
              id="flightNumber" 
              className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-4 px-4 bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500" 
              placeholder="Enter Flight Number"
              value ={formData.flightNumber}
              onChange={handleChange}
            />
          </div>
          <div>
            <label htmlFor="date" className="block text-lg font-medium text-gray-700 mb-2 mt-4">Date</label>
            <input 
              required
              type="date" 
              id="date" 
              className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-4 px-4 bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500" 
              value ={formData.date}
              onChange={handleChange}
            />
          </div>
          <div>
            <label htmlFor="aircraft" className="block text-lg font-medium text-gray-700 mb-2 mt-4">Aircraft</label>
            <select
              required
              id="aircraft"
              className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-4 px-4 bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              value={formData.aircraft}
              onChange={handleChange}
            >
              <option value="" disabled>Select Aircraft</option>
              {aircraft.map((type) => (
                <option key={type} value={type}>{type}</option>
              ))}
            </select>
          </div>
          <button 
            type="submit" 
            disabled={loading}
            className="mt-6 w-full bg-blue-600 text-white py-4 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            Generate Vouchers
          </button>
        </form>
      </div>
      {loading && <div className="text-center text-blue-600 font-medium">Processing...</div>}
      {checkResult && checkResult.exist && <div className="bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700 p-4" role="alert">
        <p className="font-bold">Notice</p>
        <p>Vouchers have already been generated for this flight number on this date.</p>
      </div>}
      {result && <div className="bg-green-100 border-l-4 border-green-500 text-green-700 p-4" role="alert">
        <p className="font-bold">Success</p>
        <p>Vouchers successfully generated, here are the seats that received vouchers:</p>
        <div className="flex justify-center mb-4 gap-2">
          {result.seats.map((seat, index) => (
            <div key={index} className="text-lg bg-emerald-200 rounded-sm p-2">{seat}</div>
          ))}
        </div>
      </div>}
    </div>
  </div>;
}
export default Page;
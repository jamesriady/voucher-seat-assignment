import { useState } from 'react';
import UserIcon from './assets/icons/user';
import IdCardIcon from './assets/icons/id_card';
import FlightIcon from './assets/icons/flight';
import AirplaneIcon from './assets/icons/airplane';
import Input from './components/Input';
import Datepicker from './components/Datepicker';
import Select from './components/Select';
import { AIRCRAFT_TYPES } from './constants/aircraft';

function App() {
  const [crewName, setCrewName] = useState('');
  const [crewId, setCrewId] = useState('');
  const [flightNumber, setFlightNumber] = useState('');
  const [flightDate, setFlightDate] = useState('');
  const [aircraftType, setAircraftType] = useState('ATR');
  const [generatedSeats, setGeneratedSeats] = useState([]);
  const [error, setError] = useState('');
  const [validationError, setValidationError] = useState({});
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setValidationError({});
    setGeneratedSeats([]);
    setLoading(true);

    try {
      const checkRes = await fetch('http://localhost:8000/api/check', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ flightNumber, date: flightDate }),
      });

      const checkData = await checkRes.json();
      if (!checkRes.ok) {
        if (checkRes.status === 400 && checkData?.errors) {
          setValidationError(checkData?.errors || {});
          setLoading(false);
          return;
        } else {
          throw new Error(checkData?.error ||'Failed to check for existing vouchers.');
        }      
      } 

      const { exists } = checkData;
      if (exists) {
        setError("Vouchers for this flight on the specified date have already been generated.");
        setLoading(false);
        return;
      }

      const generateRes = await fetch('http://localhost:8000/api/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name: crewName,
          id: crewId,
          flightNumber,
          date: flightDate,
          aircraft: aircraftType,
        }),
      });
      
      const generateData = await generateRes.json();
      if (!generateRes.ok || !generateData.success) {
        if (generateRes.status === 400 && generateData?.errors) {
          setValidationError(generateData?.errors || {});
          setLoading(false);
          return;
        } else {
          throw new Error(generateData?.error || 'Failed to generate vouchers.');
        }      
      } 

      if (generateData.success) setGeneratedSeats(generateData.seats);
    } catch (err) {
      setError(err.message || 'An unexpected error occurred.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center p-4">
      <div className="w-full max-w-4xl bg-white rounded-2xl shadow-xl flex overflow-hidden">
        <div className="hidden md:flex w-1/3 bg-gradient-to-br bg-primary p-8 flex-col justify-between text-white">
          <div>
            <h1 className="text-3xl font-bold">Airline Voucher</h1>
            <p className="mt-2 text-indigo-200">Exclusive Seat Assignments</p>
          </div>
          <div className="text-center">
            <FlightIcon />
            <p className="text-sm text-indigo-200 mt-4">Powered by modern technology for seamless airline operations.</p>
          </div>
        </div>

        <div className="w-full md:w-2/3 p-8 md:p-12">
          <h2 className="text-2xl font-bold text-gray-800 mb-2">Generate Crew Vouchers</h2>
          <p className="text-gray-500 mb-8">Enter flight details to assign random seats.</p>

          <form onSubmit={handleSubmit} className="space-y-5">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-5">
                <Input
                  icon={UserIcon} 
                  id="crew_name"
                  placeholder="Crew Name"
                  value={crewName}
                  onChange={(e) => setCrewName(e.target.value)}
                  error={validationError?.name}
                />
                <Input
                  icon={IdCardIcon} 
                  id="crew_id"
                  placeholder="Crew ID"
                  value={crewId}
                  onChange={(e) => setCrewId(e.target.value)}
                  error={validationError?.id}
                />
            </div>
            <Input
              icon={FlightIcon} 
              id="flight_number"
              placeholder="Flight Number (e.g., GA123)"
              value={flightNumber}
              onChange={(e) => setFlightNumber(e.target.value)}
              error={validationError?.flightNumber}
            />
            <div className="grid grid-cols-1 md:grid-cols-2 gap-5">
              <Datepicker
                id="date"
                value={flightDate}
                onChange={(e) => setFlightDate(e.target.value)}
                error={validationError?.date}
              />
              <Select
                icon={AirplaneIcon}
                id="aircraft_type"
                value={aircraftType} 
                onChange={(e) => setAircraftType(e.target.value)} 
                options={AIRCRAFT_TYPES}
                error={validationError?.aircraft}
              />
            </div>
            <button
                type="submit"
                disabled={loading}
                className="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-orange-500 hover:bg-orange-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:bg-blue-300 transition-colors"
            >
                {loading ? 'Generating...' : 'Generate Vouchers'}
            </button>
          </form>

          {error && (
            <div className="mt-6 p-4 bg-red-100 border-l-4 border-red-500 text-red-700 rounded-r-lg">
                <p className="font-bold">Error</p>
                <p>{error}</p>
            </div>
          )}

          {generatedSeats.length > 0 && (
            <div className="mt-6 p-4 bg-green-100 border-l-4 border-green-500 rounded-r-lg">
                <h3 className="font-bold text-green-800">Vouchers Generated Successfully!</h3>
                <p className="text-gray-700 mb-4">Assigned seats for Flight {flightNumber}:</p>
                <div className="flex justify-around text-center gap-4">
                {generatedSeats.map((seat, index) => (
                    <div key={index} className="flex-1 bg-white p-4 rounded-lg shadow-md border">
                        <p className="text-sm text-gray-500">Seat</p>
                        <p className="text-3xl font-bold text-primary">{seat}</p>
                    </div>
                ))}
                </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
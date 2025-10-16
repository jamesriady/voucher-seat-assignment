import CalendarIcon from '../assets/icons/calendar';

export default function Datepicker({ 
    error,
    ...props
}) {
    return (
        <div>
            <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"><CalendarIcon /></div>
                <input type="date" {...props} className="w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500 text-gray-500"/>
            </div>
            { error && <p className="text-red-600 text-sm mt-1">{error}</p> }
        </div>
    );
}
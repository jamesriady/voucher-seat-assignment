export default function Input({ 
    icon: Icon,
    error,
    ...props
}) {
    return (
        <div>
            <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"><Icon /></div>
                <input type="text" {...props} className="w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
            </div>
            { error && <p className="text-red-600 text-sm mt-1">{error}</p> }
        </div>
    );
}
import { CreateGuest } from './create-guest';
import { SelectGuest } from './select-guest';
import { getParty } from '../../hooks/party';
import { useParams, useNavigate } from 'react-router-dom';

export const Login = () => {
  const navigate = useNavigate()
  const { code } = useParams();
  if (code === undefined) {
      navigate('/')
  }
  const { data, isLoading, error } = getParty(code);
  
  if (isLoading) return <div className="flex justify-center p-20 animate-pulse text-gray-500">Loading party details...</div>;
  if (error || !data) return <div className="p-6 text-red-600 bg-red-50 rounded-lg m-4">Error: {error?.message || "Data missing"}</div>;

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4">
      <div className="max-w-2xl mx-auto space-y-8">
        
        {/* Event Hero Section */}
        <section className="bg-white p-8 rounded-2xl shadow-sm border border-gray-100 text-center">
          <h1 className="text-xs font-bold uppercase tracking-widest text-blue-600 mb-2">The Event</h1>
          <h2 className="text-3xl font-extrabold text-gray-900 mb-4">{data.Title}</h2>
          <p className="text-gray-600 mb-6 whitespace-pre-line text-left">{data.Description}</p>
          <div className="flex flex-wrap justify-center gap-4 text-sm text-gray-500 font-medium">
            <span>📅 {data.Date}</span>
            <span>⏰ {data.Time}</span>
            <span className="hidden sm:inline">|</span>
            <span>📍 {data.Address}</span>
          </div>
        </section>

        {/* Action Sections */}
        <div className="grid grid-cols-1 gap-8">
          <CreateGuest code={code}/>
          <div className="relative">
            <div className="absolute inset-0 flex items-center"><span className="w-full border-t border-gray-200"></span></div>
            <div className="relative flex justify-center text-xs uppercase"><span className="bg-gray-50 px-2 text-gray-400 font-bold">Or</span></div>
          </div>
          <SelectGuest code={code}/>
        </div>
      </div>
    </div>
  );
}
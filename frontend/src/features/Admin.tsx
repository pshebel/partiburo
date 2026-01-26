import { getHome } from '../hooks/home';
import { useParams, Link, useNavigate } from 'react-router-dom';

export const AdminHome = () => {
    const navigate = useNavigate();
    const { code } = useParams();

    if (code === undefined) {
        navigate('/');
    }

    const { data, isLoading, error } = getHome(code);

    // Loading State
    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <span className="text-lg font-medium animate-pulse text-gray-600">Loading Admin Dashboard...</span>
            </div>
        );
    }

    // Error State
    if (error || !data) {
        return (
            <div className="p-6 m-4 bg-red-50 border border-red-200 rounded-lg text-red-700">
                <span className="font-bold">Error:</span> {error?.message || "Failed to load event data"}
            </div>
        );
    }

    return (
        <div className="max-w-4xl mx-auto p-6 space-y-12 bg-white min-h-screen">
            
            {/* Admin Header Section */}
            <header className="flex justify-between items-start border-b pb-8">
                <div>
                    <h1 className="text-xs font-bold uppercase tracking-widest text-red-600 mb-2">Admin Dashboard</h1>
                    <h2 className="text-4xl font-extrabold text-gray-900 mb-2">{data.Title}</h2>
                    <p className="text-gray-500 font-medium">Manage your event and communicate with guests.</p>
                </div>
                <div className="flex flex-col gap-2">
                    <Link to={`/announcement/${code}`} className="bg-orange-600 text-white px-6 py-2 rounded-lg text-sm font-bold hover:bg-orange-700 transition text-center shadow-lg">
                        Post Announcement
                    </Link>
                </div>
            </header>


            {/* Guest Management */}
            <section>
                <h1 className="text-xs font-bold uppercase tracking-widest text-green-600 mb-6">Guests</h1>
                <div className="overflow-hidden border border-gray-100 rounded-xl">
                    <table className="min-w-full divide-y divide-gray-200">
                        <thead className="bg-gray-50">
                            <tr>
                                <th className="px-6 py-3 text-left text-[10px] font-bold text-gray-500 uppercase">Name</th>
                                <th className="px-6 py-3 text-left text-[10px] font-bold text-gray-500 uppercase">Status</th>
                            </tr>
                        </thead>
                        <tbody className="bg-white divide-y divide-gray-200 text-sm">
                            {data.Guests.map((g, i) => (
                                <tr key={i}>
                                    <td className="px-6 py-4 font-medium text-gray-900">{g.name}</td>
                                    <td className="px-6 py-4">
                                        <span className={`px-2 py-1 rounded-full text-[10px] font-bold ${g.status === 'going' ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-500'}`}>
                                            {g.status.toUpperCase()}
                                        </span>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            </section>
        </div>
    );
}
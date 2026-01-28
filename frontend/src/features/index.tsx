import { useState, useMemo, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { getTitles } from '../hooks/party';
import { getCodes, syncCodes } from '../hooks/identity';
import { Header } from './Header';

export const Index = () => {
    const navigate = useNavigate();
    const [joinCode, setJoinCode] = useState('');
    
    const codes = getCodes() || [];
    const req = useMemo(() => ({ codes }), [codes.length]);
    // const req: TitleRequest = { codes };
    const { data, error } = getTitles(req, codes.length > 0);
    // Sync Logic
    useEffect(() => {
        if (data?.titles) {
            const serverCodes = Object.keys(data.titles);
            
            // If server has fewer codes than we do, some were deleted
            if (serverCodes.length < codes.length) {
                syncCodes(serverCodes);
            }
        }
    }, [data, codes.length]);

    const handleJoin = (e: React.FormEvent) => {
        e.preventDefault();
        if (joinCode.trim()) {
            navigate(`/login/${joinCode.trim()}`);
        }
    };

    // if (isLoading) return <div className="p-8 text-center text-gray-500">Loading your parties...</div>;

    return (
        <div className="max-w-2xl mx-auto p-6 space-y-8">
            <Header />
            {/* Create Party Action (Moved out of Header) */}
            <section className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                <div className="flex justify-between items-center">
                    <h2 className="text-lg font-semibold">Create a party</h2>
                    <div className="flex justify-end">
                        <Link 
                            to="/party" 
                            className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg font-medium transition shadow-sm"
                        >
                            Create
                        </Link>
                    </div>
                </div>
            </section>
                
            {/* Join Section */}
            <section className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                <h2 className="text-lg font-semibold mb-4">Join a Party</h2>
                <form onSubmit={handleJoin} className="flex gap-2">
                    <input 
                        type="text"
                        placeholder="Enter party code (e.g. abc123)"
                        value={joinCode}
                        onChange={(e) => setJoinCode(e.target.value)}
                        className="flex-1 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500"
                    />
                    <button 
                        type="submit"
                        className="bg-gray-800 text-white px-6 py-2 rounded-lg hover:bg-black transition"
                    >
                        Join
                    </button>
                </form>
            </section>

            {/* List Section */}
            <section>
                <h2 className="text-gray-500 uppercase tracking-wider text-xs font-bold mb-4">Your Recent Parties</h2>
                {codes.length === 0 || error ? (
                    <p className="text-gray-400 italic">No parties found yet.</p>
                ) : (
                    <ul className="grid gap-3">
                        {data && Object.keys(data.titles).map((key) => (
                            <li key={key}>
                                <Link 
                                    to={`/${key}`} 
                                    className="block p-4 bg-white border border-gray-200 rounded-lg hover:border-indigo-500 hover:shadow-md transition group"
                                >
                                    <div className="flex justify-between items-center">
                                        <span className="font-semibold text-gray-700 group-hover:text-indigo-600">
                                            {data.titles[key]}
                                        </span>
                                        <span className="text-sm text-gray-400 font-mono uppercase">{key}</span>
                                    </div>
                                </Link>
                            </li>
                        ))}
                    </ul>
                )}
            </section>
        </div>
    );
};
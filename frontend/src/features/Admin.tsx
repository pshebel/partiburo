import { getParty } from '../hooks/party';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { PartyForm } from './PartyForm';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import {useState } from 'react'
import { Header } from './Header'

export const AdminHome = () => {
    const navigate = useNavigate()
    const { code } = useParams();
    const queryClient = useQueryClient();
    const { data, isLoading, error } = getParty(code);

    // State to track which announcement is being edited
    const [editingId, setEditingId] = useState<number | null>(null);
    const [editValues, setEditValues] = useState({ header: '', body: '' });

    // --- Mutations ---
    const updatePartyMutation = useMutation({
        mutationFn: async (values: any) => {
            const res = await fetch(`${import.meta.env.VITE_API_URL}/admin/party/${code}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(values),
            });
            if (!res.ok) throw new Error('Failed to update party');
            return res.json();
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['party', code] });
            alert("Party updated!");
        }
    });

    const deletePartyMutation = useMutation({
        mutationFn: async () => {
            const res = await fetch(`${import.meta.env.VITE_API_URL}/admin/party/${code}`, {
                method: 'DELETE',
            });
            if (!res.ok) throw new Error('Delete failed');
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['party', code] });
        }
    });

    const deleteAnnouncementMutation = useMutation({
        mutationFn: async (id: number) => {
            const res = await fetch(`${import.meta.env.VITE_API_URL}/announcement/${code}`, {
                method: 'DELETE',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({id: id}),
            });
            if (!res.ok) throw new Error('Delete failed');
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['party', code] });
        }
    });

    const updateAnnouncementMutation = useMutation({
        mutationFn: async (id: number) => {
            const body = {
                id: id,
                ...editValues
            }
            const res = await fetch(`${import.meta.env.VITE_API_URL}/announcement/${code}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(body),
            });
            if (!res.ok) throw new Error('Update failed');
        },
        onSuccess: () => {
            setEditingId(null);
            queryClient.invalidateQueries({ queryKey: ['party', code] });
        }
    });

    if (isLoading) return <div>Loading...</div>;
    if (error || !data) {
        if (error || !data) {
            return (
                <div className="min-h-screen bg-gray-50 flex items-center justify-center p-6">
                    <div className="max-w-md w-full p-6 text-center bg-white rounded-2xl shadow-sm border border-red-100">
                    <div className="text-red-500 text-4xl mb-4">⚠️</div>
                    <h2 className="text-xl font-bold text-gray-900 mb-2">Party Not Found</h2>
                    <p className="text-gray-600 mb-6">
                        {error?.message || "We couldn't find a party with that code. Please check your link and try again."}
                    </p>
                    <button 
                        onClick={() => navigate('/')}
                        className="w-full py-2 bg-gray-900 text-white rounded-lg font-medium hover:bg-gray-800 transition-colors"
                    >
                        Go Back Home
                    </button>
                    </div>
                </div>
            );
        }
    }

    return (
        <div className="max-w-4xl mx-auto p-6 space-y-12 bg-white min-h-screen">
            <Header />
            <section className="flex justify-between items-center border-b pb-8">
                <div>
                    <h1 className="text-xs font-bold uppercase tracking-widest text-red-600 mb-2">Admin Dashboard</h1>
                    <h2 className="text-4xl font-extrabold text-gray-900 mb-2">{data.title}</h2>
                    
                </div>
                <button 
                    onClick={() => {
                        if(confirm("Delete this party?")) deletePartyMutation.mutate();
                    }}
                    className="bg-red-600 text-white px-6 py-2 rounded-lg text-sm font-bold shadow-lg"
                >
                    Delete Party
                </button>
            </section>
            {/* Announcements List */}
            <section className="space-y-4">
                <div className="flex justify-between items-start">
                    <h3 className="text-xl font-bold text-gray-900">Recent Announcements</h3>
                    <Link to={`/announcement/${code}`} className="bg-orange-600 text-white px-6 py-2 rounded-lg text-sm font-bold shadow-lg">
                        Post Announcement
                    </Link>
                </div>
                <div className="grid gap-4">
                    {data.announcements.sort((a,b) => Date.parse(b.created_at) - Date.parse(a.created_at)).map((ann: any) => (
                        <div key={ann.id} className="p-4 border rounded-xl shadow-sm bg-white relative group">
                            {editingId === ann.id ? (
                                <div className="space-y-3">
                                    <input 
                                        className="w-full p-2 border rounded"
                                        value={editValues.header}
                                        onChange={e => setEditValues({...editValues, header: e.target.value})}
                                    />
                                    <textarea 
                                        className="w-full p-2 border rounded"
                                        value={editValues.body}
                                        onChange={e => setEditValues({...editValues, body: e.target.value})}
                                    />
                                    <div className="flex gap-2">
                                        <button 
                                            onClick={() => updateAnnouncementMutation.mutate(ann.id)}
                                            className="text-sm bg-green-600 text-white px-3 py-1 rounded"
                                        >
                                            Save
                                        </button>
                                        <button 
                                            onClick={() => setEditingId(null)}
                                            className="text-sm bg-gray-200 px-3 py-1 rounded"
                                        >
                                            Cancel
                                        </button>
                                    </div>
                                </div>
                            ) : (
                                <>
                                    <div className="flex justify-between items-start">
                                        <div>
                                            <h4 className="font-bold text-lg">{ann.header}</h4>
                                            <p className="text-gray-600">{ann.body}</p>
                                        </div>
                                        <div className="flex gap-2 opacity-0 group-hover:opacity-100 transition">
                                            <button 
                                                onClick={() => {
                                                    setEditingId(ann.id);
                                                    setEditValues({ header: ann.header, body: ann.body });
                                                }}
                                                className="text-blue-600 text-sm font-semibold"
                                            >
                                                Edit
                                            </button>
                                            <button 
                                                onClick={() => {
                                                    if(confirm("Delete this announcement?")) deleteAnnouncementMutation.mutate(ann.id);
                                                }}
                                                className="text-red-600 text-sm font-semibold"
                                            >
                                                Delete
                                            </button>
                                        </div>
                                    </div>
                                </>
                            )}
                        </div>
                    ))}
                    {(!data.announcements || data.announcements.length === 0) && (
                        <p className="text-gray-400 italic">No announcements posted yet.</p>
                    )}
                </div>
            </section>

            <hr />
            {/* Party Edit Form */}
            <section className="bg-gray-50 p-8 rounded-2xl border border-gray-100">
                <h3 className="text-xl font-bold text-gray-900 mb-6">Edit Party Details</h3>
                <PartyForm 
                    buttonLabel="Update Party"
                    initialValues={data}
                    onSubmit={async (values) => {
                        await updatePartyMutation.mutateAsync(values);
                    }}
                />
            </section>
        </div>
    );
}
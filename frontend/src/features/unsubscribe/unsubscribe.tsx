import { useParams, Link } from 'react-router-dom';
import { postUnsubscribe } from '../../hooks/unsubscribe'; // Assuming this hook exists based on your pattern

export const Unsubscribe = () => {
    const { email } = useParams();

    // Reusable Error View to keep the component clean
    const ErrorView = ({ message }: { message: string }) => (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
            <div className="max-w-md w-full bg-white p-8 rounded-2xl shadow-sm border border-red-100 text-center">
                <div className="w-16 h-16 bg-red-50 text-red-500 rounded-full flex items-center justify-center mx-auto mb-4 text-2xl font-bold">
                    !
                </div>
                <h1 className="text-xl font-bold text-gray-900 mb-2">Unsubscribe Failed</h1>
                <p className="text-gray-600 mb-6">{message}</p>
                <a href="mailto:support@partiburo.com" className="text-blue-600 font-medium hover:underline">
                    Contact Support
                </a>
            </div>
        </div>
    );

    if (email === undefined) {
        return <ErrorView message="We couldn't find an email to unsubscribe. Please report this to support@partiburo.com" />;
    }

    const { data, isLoading, error } = postUnsubscribe(email);

    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="flex flex-col items-center gap-4">
                    <div className="w-12 h-12 border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
                    <span className="text-gray-600 font-medium tracking-wide">Processing request...</span>
                </div>
            </div>
        );
    }

    if (error || data?.Code !== 200) {
        return <ErrorView message="Failed to unsubscribe your email. Please report this to support@partiburo.com" />;
    }

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
            <div className="max-w-md w-full bg-white p-10 rounded-2xl shadow-xl border border-gray-100 text-center animate-in fade-in zoom-in duration-300">
                <div className="w-20 h-20 bg-gray-100 text-gray-500 rounded-full flex items-center justify-center mx-auto mb-6 text-4xl">
                    ∅
                </div>
                <h1 className="text-2xl font-extrabold text-gray-900 mb-2">Unsubscribed</h1>
                <p className="text-gray-600 mb-8 leading-relaxed">
                    You have been successfully removed from our list for <span className="font-semibold text-gray-800">{email}</span>.
                </p>
                <Link 
                    to="/" 
                    className="inline-block w-full bg-gray-900 text-white font-bold py-3 px-6 rounded-xl hover:bg-black transition-colors shadow-lg shadow-gray-200"
                >
                    Back home
                </Link>
            </div>
        </div>
    );
};
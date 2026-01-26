import { useParams, Link } from 'react-router-dom';
import { createConfirm } from '../../hooks/confirm';

export const Confirm = () => {
    const { code, passcode } = useParams();

    // Helper for error messages to keep the JSX dry
    const ErrorState = ({ message }: { message: string }) => (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
            <div className="max-w-md w-full bg-white p-8 rounded-2xl shadow-sm border border-red-100 text-center">
                <div className="w-16 h-16 bg-red-100 text-red-600 rounded-full flex items-center justify-center mx-auto mb-4 text-2xl">
                    ✕
                </div>
                <h1 className="text-xl font-bold text-gray-900 mb-2">Confirmation Failed</h1>
                <p className="text-gray-600 mb-6">{message}</p>
                <a href="mailto:support@partiburo.com" className="text-blue-600 font-medium hover:underline">
                    Contact Support
                </a>
            </div>
        </div>
    );

    if (code === undefined || passcode === undefined) {
        return <ErrorState message="Invalid confirmation link. Please report this to support@partiburo.com" />;
    }

    const { data, isLoading, error } = createConfirm(code, passcode);

    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="flex flex-col items-center gap-4">
                    <div className="w-12 h-12 border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
                    <span className="text-gray-600 font-medium">Verifying your email...</span>
                </div>
            </div>
        );
    }

    if (error || data?.Code !== 200) {
        return <ErrorState message="We couldn't confirm your email. It may have expired or already been used." />;
    }

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
            <div className="max-w-md w-full bg-white p-10 rounded-2xl shadow-xl border border-gray-100 text-center animate-in fade-in zoom-in duration-300">
                <div className="w-20 h-20 bg-green-100 text-green-600 rounded-full flex items-center justify-center mx-auto mb-6 text-4xl">
                    ✓
                </div>
                <h1 className="text-2xl font-extrabold text-gray-900 mb-2">Success!</h1>
                <p className="text-gray-600 mb-8">
                    Your email has been successfully confirmed.
                </p>
                <Link 
                    to="/" 
                    className="inline-block w-full bg-blue-600 text-white font-bold py-3 px-6 rounded-xl hover:bg-blue-700 transition-colors shadow-lg shadow-blue-200"
                >
                    Return to Home
                </Link>
            </div>
        </div>
    );
};
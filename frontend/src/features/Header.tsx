import { Link } from 'react-router-dom';

interface HeaderProps {
    title?: string;
    children?: React.ReactNode;
}

export const Header = ({ title = "Partiburo", children }: HeaderProps) => {
    return (
        <header className="flex justify-between items-center border-b pb-4 mb-8">
            <Link to="/" className="hover:opacity-80 transition">
                <h1 className="text-2xl font-bold text-gray-800">{title}</h1>
            </Link>
            <div className="flex items-center gap-4">
                {children}
            </div>
        </header>
    );
};
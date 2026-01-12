import { useParams } from 'react-router-dom';
import { Link } from 'react-router-dom'
import { createConfirm } from '../../hooks/confirm';

export const Confirm = () => {
    const { email, passcode } = useParams();

    if (email == undefined || passcode == undefined) {
        return (<div>Failed to confirm your email, please report this to support@partiburo.com</div>)
    }

    const { data, isLoading, error } = createConfirm(email, passcode);

    if (isLoading) {
        return (<div>Loading</div>)
    }
    if (error) {
        return (<div>Failed to confirm your email, please report this to support@partiburo.com</div>)
    }

    if (data?.Code !== 200) {
        return (<div>Failed to confirm your email, please report this to support@partiburo.com</div>)
    }

    return (
        <div>
            <h1>Your email has been confirmed.</h1>
            <Link to="/">Back home</Link>
        </div>
    )
}
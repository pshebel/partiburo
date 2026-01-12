import { useParams } from 'react-router-dom';
import { Link } from 'react-router-dom'
import { postUnsubscribe } from '../../hooks/unsubscribe';

export const Unsubscribe = () => {
    const { email } = useParams();

    if (email == undefined) {
        return (<div>Failed to unsubscribe, please report this to support@partiburo.com</div>)
    }

    const { data, isLoading, error } = postUnsubscribe(email);

    if (isLoading) {
        return (<div>Loading</div>)
    }
    if (error) {
        return (<div>Failed to unsubscribe, please report this to support@partiburo.com</div>)
    }

    if (data?.Code !== 200) {
        return (<div>Failed to unsubscribe, please report this to support@partiburo.com</div>)
    }

    return (
        <div>
            <h1>You have been unsubscribed.</h1>
            <Link to="/">Back home</Link>
        </div>
    )
}
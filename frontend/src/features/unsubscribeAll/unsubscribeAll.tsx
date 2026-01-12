import { useParams } from 'react-router-dom';
import { Link } from 'react-router-dom'
import { postUnsubscribeAll } from '../../hooks/unsubscribe';

export const UnsubscribeAll = () => {
    const { email } = useParams();

    if (email == undefined) {
        return (<div>Failed to unsubscribe, please report this to support@partiburo.com</div>)
    }

    const { data, isLoading, error } = postUnsubscribeAll(email);

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
            <h1>You have been unsubscribed from all future communication from partiburo. To undo this, please contact support@partiburo.com</h1>
            <Link to="/">Back home</Link>
        </div>
    )
}
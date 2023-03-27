import React, {useState} from 'react';

const Bulk = () => {
    const [results, setResults] = useState([]);
    const [count, setCount] = useState([]);
    const bulk = () => {
        const socket = new WebSocket('ws://localhost:8080/ws');

        socket.addEventListener('message', (event) => {
            const data = JSON.parse(event.data);
            setResults((prevResults) => [data, ...prevResults]);
            setCount((prevResults) => [prevResults.length + 1, ...prevResults]);
        });
        const template = {
            template: "firstTemplate"
        }

        setTimeout(() => socket.send(JSON.stringify(template)), 1e3);
    }

    const renderResults = () => {
        if (results.length === 0) {
            return null;
        }

        return (
            <table>
                <thead>
                <tr>
                    <th>Order Index</th>
                    <th>Value</th>
                </tr>
                </thead>
                <tbody>
                {results.map((data, index) => (
                    <tr key={index}>
                        <td>{count[index]}</td>
                        <td>{data.email}</td>
                        <td>{data.status === 200 ? "Sent" : "Error"}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        );
    };

    return (
        <>
            <button onClick={bulk}>Bulk</button>
            {renderResults()}
        </>
    );
};

export default Bulk;

import React, {SyntheticEvent, useState} from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';
import './styles/App.scss';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

function App() {
    let graph;
    const [countryStatistic, setCountryStatistic] = useState(Array<{x: string; y: number}>());
    const [country, setCountry] = useState('');
    const [statusCode, setStatusCode] = useState(0);
    const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun",
        "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
    ];

    const getStatistic = async (e: SyntheticEvent) => {
        e.preventDefault();
        let countryResponse = Array<{x: string; y: number}>();

        const response = await fetch('https://api.covid19api.com/country/'+country+'/status/confirmed?from=2020-03-01T00:00:00Z&to=2020-04-01T00:00:00Z', {
            method: 'GET'
        })

        const responseJSON = await response.json();

        for (let i = 0; i < responseJSON.length; i++) {
            let date = new Date(responseJSON[i]["Date"])
            let dateString = (date.getDate()).toString() + "-" + monthNames[date.getMonth()]
            countryResponse.push({x: dateString, y: responseJSON[i]["Cases"]})
        }

        setCountryStatistic(countryResponse);
        setStatusCode(response.status);
    }

    if (statusCode === 200) {
        const options = {
        };

        const data = {
            datasets: [
                {
                    label: country,
                    data: countryStatistic,
                    borderColor: 'rgb(53, 162, 235)',
                    backgroundColor: 'rgba(53, 162, 235, 0.5)',
                },
            ],
        };

        graph = (
            <div className="Statistic">
                <Line options={options} data={data} />
            </div>
        )
    }

    return (
        <div className="App">
            <form onSubmit={getStatistic}>
                <input type="text" className="inputText" placeholder="Russia"
                       onChange={e => setCountry(e.target.value)}
                />

                <button className="buttonStatistic" type="submit">Enter</button>
            </form>
            {graph}
        </div>
    );
}

export default App;

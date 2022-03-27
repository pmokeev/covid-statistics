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
    ChartData,
    ScatterDataPoint,
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

type ChartDataType = ChartData<"line", (number | ScatterDataPoint | null)[], unknown> | null;

const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
];

function App() {
    const [countryName, setCountryName] = useState('');
    const [chartData, setChartData] = useState<ChartDataType>(null);

    const chartOptions = {};

    const handleSubmit = async (e: SyntheticEvent) => {
        e.preventDefault();
        
        const response = await fetch('http://localhost:8000/api/?country='+countryName, {
            method: 'GET'
        })
        
        const responseJSON = await response.json();
        if (response.status === 200) {
            let countryResponse = Array<{x: string; y: number}>();
            for (let i = 0; i < responseJSON.length; i++) {
                const date = new Date(responseJSON[i]["Date"])
                const dateString = (date.getDate()).toString() + "-" + monthNames[date.getMonth()]
                countryResponse.push({x: dateString, y: responseJSON[i]["Cases"]})
            }
            
            setChartData({
                datasets: [
                    {
                        label: countryName,
                        // @ts-ignore
                        data: countryResponse,
                        borderColor: 'rgb(53, 162, 235)',
                        backgroundColor: 'rgba(53, 162, 235, 0.5)',
                    },
                ],
            });
        } else {
            setChartData(null);
        }
    }

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => 
        setCountryName(event.target.value);

    return (
        <div className="App">
            <form onSubmit={handleSubmit}>
                <input type="text" className="inputText" placeholder="Russia" onChange={handleChange} />

                <button className="buttonStatistic" type="submit">Enter</button>
            </form>
            <div className="Statistic">
                {chartData ? <Line options={chartOptions} data={chartData} /> : <h1>No data</h1>}
            </div>
        </div>
    );
}

export default App;

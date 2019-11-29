declare var ApexCharts:any; // Magic

const socket = new WebSocket('ws://localhost:8080/socket');

var options = {
    chart: {
        type: 'radialBar',
        // offsetY: -20
    },
    plotOptions: {
        radialBar: {
            startAngle: -135,
            endAngle: 135,
            track: {
                background: "#e7e7e7",
                strokeWidth: '97%',
                margin: 5, // margin is in pixels
            },
            dataLabels: {
                name: {
                    // fontSize: '16px',
                    color: '#fff',
                    // offsetY: 44
                },   
                value: {
                    // offsetY: 76,
                    // fontSize: '22px',
                    color: '#fff',
                    formatter: function (val: any) {
                        return val + "%";
                    }
                }                     
            }
        }
    },
    series: [0],
    labels: ['TPS'],
    legend: {
        labels: {
            colors: ['#fff']
        }
    },
   
}

var chart = new ApexCharts(
    document.querySelector("#canvas"),
    options
);

chart.render();

// let lastPercents: {[key: string]: number} = {};
socket.onmessage = function (event) {
    const sensors = JSON.parse(event.data);
    
    if (sensors.tps) {
        const volts = getADS1115Voltage(sensors.tps);
        const percent: number = Math.round(volts/5.2*100);
        // console.log(volts, percent);
        
        chart.updateSeries([percent], false);
    }
}

function getADS1115Voltage(n: number): number {
    const max = 6.144;
    let volts = n/((1<<15)-1)*max;

    if (volts < 0 || volts > max) {
        volts = 0;
    }

    return volts;
}
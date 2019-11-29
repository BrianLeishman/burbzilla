"use strict";
const socket = new WebSocket('ws://localhost:8080/socket');
var options = {
    chart: {
        type: 'radialBar',
    },
    plotOptions: {
        radialBar: {
            startAngle: -135,
            endAngle: 135,
            track: {
                background: "#e7e7e7",
                strokeWidth: '97%',
                margin: 5,
            },
            dataLabels: {
                name: {
                    color: '#fff',
                },
                value: {
                    color: '#fff',
                    formatter: function (val) {
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
};
var chart = new ApexCharts(document.querySelector("#canvas"), options);
chart.render();
socket.onmessage = function (event) {
    const sensors = JSON.parse(event.data);
    if (sensors.tps) {
        const volts = getADS1115Voltage(sensors.tps);
        const percent = Math.round(volts / 5.2 * 100);
        chart.updateSeries([percent], false);
    }
};
function getADS1115Voltage(n) {
    const max = 6.144;
    let volts = n / ((1 << 15) - 1) * max;
    if (volts < 0 || volts > max) {
        volts = 0;
    }
    return volts;
}
//# sourceMappingURL=main.js.map
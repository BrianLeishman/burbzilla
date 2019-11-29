"use strict";
const socket = new WebSocket('ws://localhost:8080/socket');
let lastValues = {};
let charts = {};
socket.onmessage = function (event) {
    const sensors = JSON.parse(event.data);
    const names = Object.keys(sensors);
    for (let i = 0; i < names.length; i++) {
        const name = names[i];
        if (!(name in lastValues)) {
            lastValues[name] = 0;
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
                        },
                        dataLabels: {
                            name: {
                                color: '#fff',
                            },
                            value: {
                                color: '#fff',
                                formatter: function (val) {
                                    return Math.round(val) + "%";
                                }
                            }
                        }
                    }
                },
                series: [0],
                labels: ['Sensor'],
            };
            let sensorUnknown = false;
            switch (name.toLowerCase()) {
                case 'tps':
                    options.labels = ['TPS'];
                    break;
                case 'oil':
                    options.labels = ['Oil'];
                    options.plotOptions.radialBar.dataLabels.value.formatter = function (val) {
                        return Math.round(val * 1.5) + " PSI";
                    };
                    break;
                default:
                    sensorUnknown = true;
            }
            if (!sensorUnknown) {
                charts[name] = new ApexCharts(document.getElementById(name), options);
                charts[name].render();
            }
        }
    }
    if (sensors.tps) {
        const volts = getADS1115Voltage(sensors.tps);
        const percent = (volts / 5.2 * 100);
        console.log(volts);
        if (Math.abs(volts - lastValues.tps) > .05) {
            charts.tps.updateSeries([percent], false);
            lastValues.tps = volts;
        }
    }
    if (sensors.oil) {
        const volts = getADS1115Voltage(sensors.oil);
        const oldMin = .5, oldMax = 4.5;
        const newMin = 0, newMax = 100;
        const oldRange = oldMax - oldMin;
        const newRange = newMax - newMin;
        let psi = (((volts - oldMin) * newRange) / oldRange) + newMin;
        if (psi < 0) {
            psi = 0;
        }
        if (Math.abs(volts - lastValues.oil) > .005) {
            charts.oil.updateSeries([psi], false);
            lastValues.oil = volts;
        }
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
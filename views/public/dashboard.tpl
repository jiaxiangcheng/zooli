<div class="six column centered row">
    <div class="column">
        <div class="ui raised compact segment">
            <div class="title">
                <i class="user icon"></i>
                {{i18n .Lang "public_dashoard.total_customers"}}
            </div>
            <div class="ui statistic">
                <div class="value" style="color: #6495ED;">
                    {{.customercount}}
                </div>
            </div>
            <div class="ui divider"></div>
            <div class="count">
                <i class="tiny angle up icon"></i>
                <label style="font-size: 10px;">5% From last Week</label>
            </div>
        </div>
    </div>

    <div class="column">
        <div class="ui raised compact segment">
            <div class="title">
                <i class="world icon"></i>
                {{i18n .Lang "public_dashoard.total_orders"}}
            </div>
            <div class="ui statistic">
                <div class="value" style="color: #6495ED;">
                    {{.ordercount}}
                </div>
            </div>
            <div class="ui divider"></div>
            <div class="count">
                <i class="tiny angle up icon"></i>
                <label style="font-size: 10px;">3% From last Week</label>
            </div>
        </div>
    </div>

    <div class="column">
        <div class="ui raised compact segment">
            <div class="title">
                <i class="cubes icon"></i>
                {{i18n .Lang "public_dashoard.total_XXXX"}}
            </div>
            <div class="ui statistic">
                <div class="value" style="color: #6495ED;">
                    {{.servicecount}}
                </div>
            </div>
            <div class="ui divider"></div>
            <div class="count">
                <i class="tiny angle up icon"></i>
                <label style="font-size: 10px;">5% From last Week</label>
            </div>
        </div>
    </div>

    <div class="column">
        <div class="ui raised compact segment">
            <div class="title">
                <i class="shopping bag icon"></i>
                {{i18n .Lang "public_dashoard.total_XXXX"}}
            </div>
            <div class="ui statistic">
                <div class="value" style="color: #6495ED;">
                    {{.storecount}}
                </div>
            </div>
            <div class="ui divider"></div>
            <div class="count">
                <i class="tiny angle down icon"></i>
                <label style="font-size: 10px;">10% From last Week</label>
            </div>
        </div>
    </div>
</div>

<div class="two wide column">
    <div class="ui raised compact segment">
        <p style="font-size: 15px">In service</p>
        <div class="chart-container" style="position: relative; height: 20vh; width: 20vh">
            <canvas id="in_service"></canvas>
        </div>
    </div>
</div>
<div class="two wide column">
    <div class="ui raised compact segment">
        <p style="font-size: 15px">End service</p>
        <div class="chart-container" style="position: relative; height: 20vh; width: 20vh">
            <canvas id="end_service"></canvas>
        </div>
    </div>
</div>
<div class="two wide column">
    <div class="ui raised compact segment">
        <p style="font-size: 15px">Wait for payment</p>
        <div class="chart-container" style="position: relative; height: 20vh; width: 20vh">
            <canvas id="w_for_payment"></canvas>
        </div>
    </div>
</div>
<div class="two wide column">
    <div class="ui raised compact segment">
        <p style="font-size: 15px">Finished</p>
        <div class="chart-container" style="position: relative; height: 20vh; width: 20vh">
            <canvas id="finished"></canvas>
        </div>
    </div>
</div>
<div class="two wide column">
    <div class="ui raised compact segment">
        <p style="font-size: 15px">Canceled</p>
        <div class="chart-container" style="position: relative; height: 20vh; width: 20vh">
            <canvas id="canceled"></canvas>
        </div>
    </div>
</div>

<div class="sixteen wide column">
    <div class="ui raised compact segment">
        <p style="font-size: 15px">Income of products</p>
        <div class="chart-container" style="position: relative; height: 60vh; width: 50vw">
            <canvas id="income"></canvas>
        </div>
    </div>
</div>




<script>
    var ctx = document.getElementById('in_service').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'pie',

    // The data for our dataset
    data: {
        labels: ["In service", "Total"],
        datasets: [{
            label: "Finished orders",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [55, 140],
        }]
    },

    // Configuration options go here
    options: {}
    });

    var ctx = document.getElementById('end_service').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'pie',

    // The data for our dataset
    data: {
        labels: ["End service", "Total"],
        datasets: [{
            label: "Finished orders",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [62, 40],
        }]
    },

    // Configuration options go here
    options: {}
    });
    var ctx = document.getElementById('w_for_payment').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'pie',

    // The data for our dataset
    data: {
        labels: ["Wait for payment", "Total"],
        datasets: [{
            label: "Finished orders",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [97, 140],
        }]
    },

    // Configuration options go here
    options: {}
    });
    var ctx = document.getElementById('finished').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'pie',

    // The data for our dataset
    data: {
        labels: ["Finished", "Total"],
        datasets: [{
            label: "Finished orders",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [53, 140],
        }]
    },

    // Configuration options go here
    options: {}
    });
    var ctx = document.getElementById('canceled').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'pie',

    // The data for our dataset
    data: {
        labels: ["Canceled", "Total"],
        datasets: [{
            label: "Finished orders",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [20, 140],
        }]
    },

    // Configuration options go here
    options: {}
    });

    var ctx = document.getElementById('income').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
        labels: ["January", "February", "March", "April", "May", "June", "July"],
        datasets: [
            {
                label: "Income history",
                backgroundColor: 'rgb(255, 99, 132)',
                borderColor: 'rgb(255, 99, 132)',
                data: [0, 10, 5, 2, 20, 30, 45],
            },
            {
                label: "Income history2",
                backgroundColor: 'rgb(0, 255, 132)',
                borderColor: 'rgb(0, 255, 132)',
                data: [0, 5, 5, 2, 20, 30, 45],
            },
            {
                label: "Income history3",
                backgroundColor: 'rgb(0, 255, 255)',
                borderColor: 'rgb(0, 255, 255)',
                data: [10, 5, 5, 2, -10, -20, 45],
            }
        ]
    },

    // Configuration options go here
    options: {
        responsive: true,
                tooltips: {
                    mode: 'index',
                },
                hover: {
                    mode: 'index'
                },
                scales: {
                    xAxes: [{
                        scaleLabel: {
                            display: true,
                            labelString: 'Month'
                        }
                    }],
                    yAxes: [{
                        stacked: true,
                        scaleLabel: {
                            display: true,
                            labelString: 'Value'
                        }
                    }]
                }
    }
    });


</script>

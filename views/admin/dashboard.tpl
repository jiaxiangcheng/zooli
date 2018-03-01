
    <div class="six column centered row">
        <div class="column">
            <div class="ui raised compact segment">
                <div class="title">
                    <i class="user icon"></i>
                    Total Users
                </div>
                <div class="ui statistic">
                    <div class="value" style="color: #6495ED;">
                        {{.usercount}}
                    </div>
                    <div class="label">
                        Users
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
                    Total Companies
                </div>
                <div class="ui statistic">
                    <div class="value" style="color: #6495ED;">
                        {{.companycount}}
                    </div>
                    <div class="label">
                        Companies
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
                    Total Services
                </div>
                <div class="ui statistic">
                    <div class="value" style="color: #6495ED;">
                        {{.servicecount}}
                    </div>
                    <div class="label">
                        Services
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
                    Total Stores
                </div>
                <div class="ui statistic">
                    <div class="value" style="color: #6495ED;">
                        {{.storecount}}
                    </div>
                    <div class="label">
                        Stores
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




    <div class="eight wide column">
        <div class="ui raised compact segment">
                <div class="chart-container" style="height:300px; width:500px">
                    <canvas id="most_used_services"></canvas>
                </div>
        </div>
    </div>
    <div class="eight wide column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:300px; width:500px">
                <canvas id="most_used_services1"></canvas>
            </div>
        </div>
    </div>




    <div class="eight wide column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:300px; width:500px;">
                <canvas id="income"></canvas>
            </div>
        </div>
    </div>

    <div class="four wide column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:300px; width:500px;">
                <canvas id="user_activity"></canvas>
            </div>
        </div>
    </div>

    <div class="six wide column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:200px; width:200px;">
                <canvas id="finished_order"></canvas>
            </div>
        </div>
    </div>

    <div class="six wide column">
        <div class="ui raised segment">
            <div class="ui big header">Map will be here</div>
        </div>
    </div>



<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.4.0/Chart.min.js"></script>
<script src="static\js\custom-todo.js"></script>


<script>
    var ctx = document.getElementById('finished_order').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'pie',

    // The data for our dataset
    data: {
        labels: ["Finished", "Pending"],
        datasets: [{
            label: "Finished orders",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [50, 40],
        }]
    },

    // Configuration options go here
    options: {}
    });

    var ctx = document.getElementById('user_activity').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
        labels: ["Monday", "Tuesday", "Wednesdays", "Thursday", "Friday", "Saturday", "Sunday"],
        datasets: [{
            label: "Users activities",
            backgroundColor: ["rgb(255, 99, 132)", "rgb(221, 221, 221)"],
            borderColor: 'rgb(255, 99, 132)',
            data: [50, 40, 30, 100, 124, 78, 37],
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



    var ctx = document.getElementById('most_used_services').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'bar',

    // The data for our dataset
    data: {
        labels: ["January", "February", "March", "April", "May", "June", "July"],
        datasets: [{
            label: "Most used services",
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: [0, 10, 5, 2, 20, 30, 45],
        }]
    },

    // Configuration options go here
    options: {}
    });



    var ctx = document.getElementById('most_used_services1').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'bar',

    // The data for our dataset
    data: {
        labels: ["January", "February", "March", "April", "May", "June", "July"],
        datasets: [{
            label: "Most used services",
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: [0, 10, 5, 2, 20, 30, 45],
        }]
    },

    // Configuration options go here
    options: {}
    });

</script>
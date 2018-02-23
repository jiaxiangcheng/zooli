<head>
    {{template "common/flash.tpl" .}}
    {{template "common/modal.tpl" .}}
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.4.0/Chart.min.js"></script>
    <script src="static\js\custom-todo.js"></script>
</head>

<body>
    <div class="ui grid">
        <div class="six column centered row">
            <div class="column">
                <div class="title">
                    <i class="user icon"></i>
                    Total Users
                </div>
                <div class="ui statistic">
                    <div class="value">
                        {{.usercount}}
                    </div>
                    <div class="label">
                        Users
                    </div>
                </div>
                <div class="count">
                    <i class="tiny angle up icon"></i>
                    <label style="font-size: 10px;">3% From last Week</label>
                </div>
            </div>
            <div class="column">
                <div class="title">
                    <i class="world icon"></i>
                    Total Companies
                </div>
                <div class="ui statistic">
                    <div class="value">
                        {{.companycount}}
                    </div>
                    <div class="label">
                        Companies
                    </div>
                </div>
            </div>
            <div class="column">
                <div class="title">
                    <i class="cubes icon"></i>
                    Total Services
                </div>
                <div class="ui statistic">
                    <div class="value">
                        {{.servicecount}}
                    </div>
                    <div class="label">
                        Services
                    </div>
                </div>
            </div>
            <div class="column">
                <div class="title">
                    <i class="shopping bag icon"></i>
                    Total Stores
                </div>
                <div class="ui statistic">
                    <div class="value">
                        {{.storecount}}
                    </div>
                    <div class="label">
                        Stores
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="ui divider"></div>
    <div class="ui two column grid">
        <div class="column">
            <div class="chart-container" style="position: relative; height:400px; width:500px">
                <canvas id="income"></canvas>
            </div>
        </div>
        <div class="column">
            <div class="chart-container" style="position: relative; height:400px; width:500px">
                <canvas id="most_used_services"></canvas>
            </div>
        </div>
    </div>

</body>

<script>
    var ctx = document.getElementById('income').getContext('2d');
    var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
        labels: ["January", "February", "March", "April", "May", "June", "July"],
        datasets: [{
            label: "Income history",
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: [0, 10, 5, 2, 20, 30, 45],
        }]
    },

    // Configuration options go here
    options: {}
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

</script>

<style>

</style>

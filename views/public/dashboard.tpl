<h1>PUBLIC</h1>
<div class="ui grid">
    <div class="six column centered row">
        <div class="column">
            <div class="title">
                <i class="user icon"></i>
                Total Customers
            </div>
            <div class="ui statistic">
                <div class="value">
                    {{.usercount}}
                </div>
                <div class="label">
                    Customer
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
                Total orders
            </div>
            <div class="ui statistic">
                <div class="value">
                    {{.companycount}}
                </div>
                <div class="label">
                    Order
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

<div class="ui five column grid">
    <div class="column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:200px; width:200px;">
                <canvas id="in_service"></canvas>
            </div>
        </div>
    </div>
    <div class="column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:200px; width:200px;">
                <canvas id="end_service"></canvas>
            </div>
        </div>
    </div>
    <div class="column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:200px; width:200px;">
                <canvas id="w_for_payment"></canvas>
            </div>
        </div>
    </div>
    <div class="column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:200px; width:200px;">
                <canvas id="finished"></canvas>
            </div>
        </div>
    </div>
    <div class="column">
        <div class="ui raised compact segment">
            <div class="chart-container" style="height:200px; width:200px;">
                <canvas id="canceled"></canvas>
            </div>
        </div>
    </div>

</div>

<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.4.0/Chart.min.js"></script>
<script src="static\js\custom-todo.js"></script>


<script>
    var ctx = document.getElementById('in_service').getContext('2d');

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

</script>


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Sign-Up/Login Form</title>

    <link rel='stylesheet' href='./static/css/css.css' type='text/css'>
    <link rel="stylesheet" href="./static/css/style.css" type='text/css'>
    <link rel="stylesheet" href="./static/css/normalize.min.css" type='text/css'>

</head>

<body>
<div class="form">

    <ul class="tab-group">
        <li class="tab active"><a href="#signup">Sign Up</a></li>
        <li class="tab"><a href="#login">Log In</a></li>
    </ul>

    <div class="tab-content">

        <div id="signup">
            <h1>Sign Up for Student</h1>

            <form  method="POST">

                <div class="field-wrap">
                    <label>
                        Student ID<span class="req">*</span>
                    </label>
                    <input name="sid" type="text" required autocomplete="off"/>
                </div>

                <div class="field-wrap">
                    <label>
                        Student name<span class="req">*</span>
                    </label>
                    <input name="sname" type="text" required autocomplete="off"/>
                </div>

                <div class="field-wrap">
                    <label>
                        Input the Password<span class="req">*</span>
                    </label>
                    <input name="pwd" type="password" required autocomplete="off"/>
                </div>

                <button type="submit" class="button button-block"/>
                Get Started</button>

            </form>

        </div>

        <div id="login">
            <h1>Student user, welcome Back!</h1>

            <form method="POST">

                <div class="field-wrap">
                    <label>
                        Student ID<span class="req">*</span>
                    </label>
                    <input name="sid" type="text" required autocomplete="off"/>
                </div>

                <div class="field-wrap">
                    <label>
                        Input Your Password<span class="req">*</span>
                    </label>
                    <input name="pwdin" type="password" required autocomplete="off"/>
                </div>

                <button class="button button-block"/>
                Log In</button>

            </form>

        </div>

    </div><!-- tab-content -->

</div> <!-- /form -->
<script src='./static/js/jquery.min.js'></script>

<script src="./static/js/index.js"></script>


</body>
</html>

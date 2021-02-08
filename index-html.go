package main

var indexHtml = `
<!doctype html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Random Password Please</title>
	<style type="text/css">
		body {
			font-size: 18px;
		}
		.slider {
			width: 50%;
		}
	</style>
</head>
<body>
	<div style="text-align: center">
		<p>Your random password is:</p>
		<h1 id="password">{{.Password}}</h1>
		<input type="range" min="8" max="30" value="12" class="slider" id="slider">
		<p><span id="length-label">12</span> characters</p>
		<button id="button">Another Password Please</button>
		<p><span id="counter">{{.Counter}}</span> passwords generated</p>
		<p>
				<a href="https://github.com/jbarham/random-password-please">Source</a> | <attr title="{{.Host}}/password.txt?len=n where n = 8-30">API</attr>
		</p>
	</div>
	<script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
	<script type="text/javascript">
		$(document).ready(function() {
			function getNewPassword() {
				/* Load new password via API. */
				$('#password').load('password.txt?len=' + $('#slider').val());
				$('#counter').load('counter');
			};

			$('#slider').on("input", function(event) {
				var val = $(event.target).val();
				$('#length-label').html(val);
			});

			$('#slider').change(function(event) {
				var val = $(event.target).val();
				$('#length-label').html(val);
				getNewPassword();
			});

			$('#button').click(function(event) {
				event.preventDefault();
				getNewPassword();
			});
		});
	</script>
</body>
</html>
`

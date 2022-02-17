$( function () {
	$.ajax({
		url: '/wallet',
		type: 'POST',
		success: function (response) {
			$('#public_key').text(response.public_key);
			$('#private_key').text(response.private_key);
			$('#blockchain_address').text(response.blockchain_address);
			console.info(response);
		},
		error: function (error) {
			console.error(error);
		}
	});

	$('#send_money_button').click(function () {
		let confirm_text = 'Are you sure you want to send money?';
		let confirm_result = confirm(confirm_text);
		if (confirm_result !== true) {
			alert('Operation canceled.');
			return;
		}
		let transaction_data = {
			'sender_private_key': $('#private_key').val(),
			'sender_blockchain_address': $('#blockchain_address').val(),
			'recipient_blockchain_address': $('#recipient_blockchain_address').val(),
			'sender_public_key': $('#public_key').val(),
			'value': $('#send_amount').val()
		};
		$.ajax({
			url: '/transaction',
			type: 'POST',
			contentType: 'application/json',
			data: JSON.stringify(transaction_data),
			success: function (response) {
				console.info(response);
				if (response.message === 'failed') {
					alert("failed");
					return
				}
				alert('Transaction posted successfully.');
			},
			error: function (error) {
			console.error(error);
			alert('Error posting transaction.');
			}
		});
	});

	function reload_amount() {
		let data = {'blockchain_address': $('#blockchain_address').val()};
		$.ajax({
			url: '/wallet/amount',
			type: 'GET',
			data: data,
			success: function (response) {
				console.info(response);
				let amount = response['amount'];
				$('#wallet_amount').text(amount);
			},
			error: function (error) {
				console.error(error);
			}
		
		});
	}

	/*
	$('#reload_wallet').click(function () {
		reload_amount();
	});
	*/
	

	setInterval(reload_amount, 3000);
})
/*
blockchain_server_0    | Blockchain Server:2022/02/17 01:44:21 private_key 6c9802fc269490930210850bcaafec1c346d11f90d2018a91a2b48c2896a9e7a
blockchain_server_0    | Blockchain Server:2022/02/17 01:44:21 public_key f6d101757a1ca9776ae1385efa118bc6322c4392c155d63114206a76249bc9c03e8d1569283cc002508761fc66af334bdb884e7f173f6436219dab108f0fa2f0
blockchain_server_0    | Blockchain Server:2022/02/17 01:44:21 blockchain_address 1CMzGxmFXRjwjgfUnXq51RtC6kxRv4S6mZ
blockchain_server_0    | HOSTNAME: 8151c95ed4f9
*/
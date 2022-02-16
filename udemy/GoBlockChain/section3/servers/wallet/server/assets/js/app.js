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
			'sender_public_key': $('#public_key').text(),
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
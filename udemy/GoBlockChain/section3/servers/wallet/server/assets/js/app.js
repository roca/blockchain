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
blockchain_server_0  | Blockchain Server:2022/02/17 14:58:08 private_key 0a8b595a05680ee3f29a41d387bd79a114ce62ca9b902b58aca71289822ef72d
blockchain_server_0  | Blockchain Server:2022/02/17 14:58:08 public_key 7fbf72bb3806d731fb8f201a56931704283502d2a8b3d6918f2d163becf7c10901fea15d6787aa6bef39a22804448326dac4d9fabc11aeb99807e9334f18ae8d
blockchain_server_2  | Blockchain Server:2022/02/17 14:58:08 Resolve conflicts replaced chain with longest chain
blockchain_server_0  | Blockchain Server:2022/02/17 14:58:08 blockchain_address 1MpF84ZHxjKvQddHfQ6qWwrUnT4QTdo5XP
blockchain_server_0  | HOSTNAME: 61b29b57636a
*/
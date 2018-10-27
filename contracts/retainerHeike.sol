pragma solidity ^0.4.23;

import "./helpers/Ownable.sol";
import "./helpers/ERC20.sol";

contract HeikeRetainer is Ownable {

    struct Transaction {
          address contract_;
          address to_;
          uint amount_;
          bool failed_;
        }

    mapping(address => uint[]) public transactionIndexesToSender;

    Transaction[] public transactions;

    address public owner;

    mapping(bytes32 => address) public tokens;

    ERC20 public ERC20Interface;

    event TransactionSuccessful(address indexed from_, address indexed to_, uint256 amount_);

    event TransactionFailed(address indexed from_, address indexed to_, uint256 amount_);

    constructor() public payable {
        owner = msg.sender;
    }

    function addNewToken(bytes32 symbol_, address address_) public onlyOwner returns (bool) {
        tokens[symbol_] = address_;

        return true;
    }

    function removeToken(bytes32 symbol_) public onlyOwner returns (bool) {
        require(tokens[symbol_] != 0x0);

        delete(tokens[symbol_]);

        return true;
    }

    function fundRetainer(bytes32 symbol_, uint256 amount_) public {
        require(tokens[symbol_] != 0x0);
        require(amount_ > 0);

        address contract_ = tokens[symbol_];
        address from_ = msg.sender;
        address to_ = address(this);

        ERC20Interface =  ERC20(contract_);

        uint256 transactionId = transactions.push(
            Transaction({
            contract_:  contract_,
            to_: to_,
            amount_: amount_,
            failed_: true
            })
        );

        transactionIndexesToSender[from_].push(transactionId - 1);

        if(amount_ > ERC20Interface.allowance(from_, address(this))) {
            emit TransactionFailed(from_, to_, amount_);
            revert();
        }

        ERC20Interface.transferFrom(from_, to_, amount_);

        transactions[transactionId - 1].failed_ = false;

        emit TransactionSuccessful(from_, to_, amount_);
    }

    function withdrawRetainer(bytes32 symbol_, address to_, uint256 amount_) public{

        address contract_ = tokens[symbol_];
        ERC20Interface =  ERC20(contract_);

        uint256 transactionId = transactions.push(
            Transaction({
            contract_:  contract_,
            to_: to_,
            amount_: amount_,
            failed_: true
            })
        );

       ERC20Interface.transfer(to_, amount_);

       transactions[transactionId - 1].failed_ = false;
       emit TransactionSuccessful(address(this), to_, amount_);

    }

    function() public payable {}

    function withdraw(address beneficiary) public payable onlyOwner {
        beneficiary.transfer(address(this).balance);
    }

 }


// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

contract Voting {

    mapping(address=>uint256) public  Votes;
    address[] public keyList;

    function Vote(address target) public  {
        if (Votes[target] == 0) {
            keyList.push(target);
            Votes[target] = 0;
        }
        Votes[target] += 1;
    }

    function GetVotes(address target) public view returns (uint256){
        return Votes[target];
    }

    function ResetVotes() public {
        for (uint256 i=0; i<keyList.length; i++) 
        {
            address target = keyList[i];
            delete (Votes[target]);
        }
        delete keyList;
    }
// ✅ 反转字符串 (Reverse String)
// 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
    function ReversStr(string memory str) public pure returns (string memory) {
        bytes memory b = bytes(str);
        uint256 len = b.length;
        bytes1 temp = 0;
        for( uint256 i=0; i<b.length/2; i++) {
            temp = b[i];
            b[i] = b[len-1-i];
            b[len - 1-i] = temp;
        }
        return string(b);
    }
//✅  用 solidity 实现整数转罗马数字
    function NumberToRoma(uint256 nub) public pure returns (string memory) {
        if (nub < 0 || nub >3999){
            return "";
        }
        string memory out;
        string[7] memory romaStr = ["M","D","C","L","X","V","I"];
        uint16[7] memory romaNum = [1000, 500, 100, 50, 10, 5, 1];
        for (uint8 i=0 ; i <7; i++) 
        {
            uint8 times = uint8(nub/romaNum[i]);
            if (times>0 ){
                if (romaNum[i] == 1000 || romaNum[i] == 100 || romaNum[i] == 10 || romaNum[i] == 1)
                {
                    if(times == 4){
                        out = string(abi.encodePacked(out, romaStr[i], romaStr[i-1]));
                    }else{
                        for (uint8 j=0; j<times; j++) 
                        {
                            out = string(abi.encodePacked(out, romaStr[i]));
                        }
                    }
                    nub -= romaNum[i]*times;
                }else{
                    uint8 nttimes = uint8(nub/romaNum[i+1]);
                    if (nttimes == 9){
                        out = string(abi.encodePacked(out, romaStr[i+1], romaStr[i-1]));
                        
                        nub -= romaNum[i+1]*nttimes;
                    }else{
                        for (uint8 j=0; j<times; j++) 
                        {
                            out = string(abi.encodePacked(out, romaStr[i]));
                        }
                        
                        nub -= romaNum[i]*times;
                    }

                }
            }
        }
        return  out;
    }

    function getRomaNum(bytes1 s) internal pure returns (int32){
        string[7] memory romaStr = ["M","D","C","L","X","V","I"];
        uint16[7] memory romaNum = [1000, 500, 100, 50, 10, 5, 1];
        
        for (uint8 j=0; j <7; j++) 
        {   
            bytes memory bs = bytes(romaStr[j]);
            bytes1 sb = bs[0];
            if(sb == s){
                return int16(romaNum[j]);
            }
        }
        return 0;
    }
//✅  用 solidity 实现罗马数字转数整数
    function RomaToNum(string memory roma) public pure returns(uint256){
        bytes memory bts = bytes(roma);
        uint8 len = uint8(bts.length);
        int256 out = 0;
        for (uint8 i=0; i<len; i++){
            bytes1 k = bts[i];
            bytes1 kn = "";
            int32 k_num = getRomaNum(k);
            int32 kn_num =0;
            if (i <len-1){
                kn = bts[i+1];
                kn_num = getRomaNum(kn);
            }
            if (k_num >= kn_num){
                out += k_num;
            }else{
                out-=k_num;
            }
        }
        return uint256(out);
    }
	//✅  合并两个有序数组 (Merge Sorted Array)
	function MergeSortedArray(uint256[] memory nums1, uint256[] memory nums2) public pure returns (uint256[] memory) {
		uint256 x  = 0;
		uint256 y = 0;
		uint256[] memory out = new uint256[](nums1.length + nums2.length);
		for (uint256 i = 0; i < nums1.length + nums2.length; i++) {
			if (x < nums1.length && y < nums2.length) {
				if (nums1[x] < nums2[y]) {
					out[i] = nums1[x];
					x++;
				} else {
					out[i] = nums2[y];
					y++;
				}
			} else if (x < nums1.length) {
				out[i] = nums1[x];
				x++;
			} else if (y < nums2.length) {
				out[i] = nums2[y];
				y++;
			}
		}
		return out;
	}
	//✅  二分查找 (Binary Search)
	function BinarySearch(uint256[] memory nums, uint256 target) public pure returns (bool) {
		uint256 left  = 0;
		uint256 right = nums.length -1;
		while (left <= right) {
			uint256 mid = left + (right - left) / 2;
			if (nums[mid] == target) {
				return true;
			} else if (nums[mid] < target) {
				left = mid + 1;
			} else {
				right = mid - 1;
			}
		}
		return false;
	}

}


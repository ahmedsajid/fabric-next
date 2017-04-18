/*
Copyright DTCC, IBM 2016, 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

         http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package example;

import static java.lang.String.format;
import static org.hyperledger.java.shim.ChaincodeHelper.newBadRequestResponse;

import org.hyperledger.fabric.protos.peer.ProposalResponsePackage.Response;
import org.hyperledger.java.shim.ChaincodeBase;
import org.hyperledger.java.shim.ChaincodeStub;

public class Example06 extends ChaincodeBase {
	
	@Override
	public Response init(ChaincodeStub stub) {
		return invoke(stub);
	}
	
	@Override
	public Response invoke(ChaincodeStub stub) {
		if(stub.getArgsAsStrings().isEmpty())
			return newBadRequestResponse(format("No arguments specified."));
		
		switch (stub.getArgsAsStrings().get(0)) {
		case "runtimeException":
			throw new RuntimeException("Exception thrown as requested.");
		default:
			return newBadRequestResponse(format("Invalid arguments specified"));
		}
	}
	
	@Override
	public String getChaincodeID() {
		return "Example06";
	}

	public static void main(String[] args) throws Exception {
		new Example06().start(args);
	}

}
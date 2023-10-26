package ufg.inf.cc.pfc1.blockchain.hyperledgerfabric.sdkhyperledgerfabricjava;

import java.io.FileInputStream;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.ProposalResponse;
import org.hyperledger.fabric.sdk.TransactionProposalRequest;
import org.hyperledger.fabric.sdk.ChaincodeID;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;

import java.util.Collection;
import java.util.concurrent.ExecutionException;

public class FabricClient {
    public static void main(String[] args)
            throws InvalidArgumentException,
            ProposalException,
            ExecutionException, InterruptedException,
            Exception {
        // Create a new HFClient instance 
        HFClient client = HFClient.createNewInstance();

        // Set the CryptoSuite for the client
        client.setCryptoSuite(CryptoSuite.Factory.getCryptoSuite());

        // Set the username and affiliation for the client 
        client.setUserContext(
                new MyUser(
                        "matheus",
                        "BR",
                        "Org1MSP",
                        Utils.getX509CertificateFromFile(
                            "/home/matheuslazaro/Desktop/PFC1/hyperledgerFabricTutorialStudy/SDKHyperledgerFabricJava/src/main/java/ufg/inf/cc/pfc1/blockchain/hyperledgerfabric/sdkhyperledgerfabricjava/cert.pem"
                        )
                )
        );

        // Create a new channel 
        Channel channel = client.newChannel("mychannel");

        // Add the peers to the channel 
        channel.addPeer(client.newPeer (
                "peer0.org1.example.com",
                "grpc://localhost:7051"
            )
        );

        channel.addPeer(
                client.newPeer(
                        "peer1.org1.example.com",
                        "grpc://localhost:7056"
                )
        );

        // Initialize the channel 
        channel.initialize();

        // Create a new transaction request 
        TransactionProposalRequest request
                = client.newTransactionProposalRequest();
        ChaincodeID ccid = ChaincodeID.newBuilder().setName("mycc").build();
        request.setChaincodeID(ccid);
        request.setFcn("invoke");
        request.setArgs(new String[]{"a", "b", "1"});

        // Submit the transaction request to the channel
        Collection<ProposalResponse> responses;
        responses = channel.sendTransactionProposal(
                request
        );

        // Print the endorsement results
        for (ProposalResponse response : responses) {
            System.out.println(
                    "Endorsement status: "
                    + response.getStatus()
            );

            System.out.println(
                    "Endorsement message: "
                    + response.getMessage()
            );
        }

        try {
            channel.sendTransaction(responses).get();
            System.out.println(
                    "Transaction sent to orderer."
            );
        } catch (Exception e) {
            System.out.println(
                    "Something went wrong."
            );
            e.printStackTrace();
        }
    }

    public static class Utils {

      public static X509Certificate getX509CertificateFromFile(String filename) throws Exception {
          CertificateFactory fact = CertificateFactory.getInstance("X.509");
          FileInputStream is = new FileInputStream(filename);
          X509Certificate cer = (X509Certificate) fact.generateCertificate(is);
          return cer;
      }
    }
}
